package net

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"strings"
	"sync"
	"time"

	"github.com/supernet/common/utils/bigendian"
	"github.com/supernet/gateway/enum/msg"
	protoNum "github.com/supernet/gateway/enum/msg"
	entity "github.com/supernet/gateway/internal/model/entity"
	gwpb "github.com/supernet/gateway/internal/net/gateway/pb"

	"github.com/supernet/gateway/internal/service"
	"github.com/supernet/gateway/pkg/log"
	"github.com/supernet/gateway/pkg/mysql"
	"github.com/supernet/gateway/pkg/pack"
	"github.com/supernet/gateway/pkg/redislib"
	"github.com/supernet/gateway/pkg/viper"
)

type Client struct {
	CManager *clientManager
	Conn     *WsConnection   // websocket 链接信息
	User     entity.Users    // 用户信息
	UserInfo entity.UserInfo // 用户游戏详情
	sid      string
	open     bool //是否在线
	cpack    pack.CPackage
	spack    pack.SPackage

	ip     string
	rwlock sync.RWMutex

	pingTime int64
	// webscoket长连接id
	WsConnId string

	currGrpcConn *GrpcClientConn
	// grpc stream connect id 一个客户端有多个grpcstream连接 一个grpc 流也可以供多个client 连接复用
	grpcConnMap map[string]*GrpcClientConn

	//用戶當前玩的游戲
	currentGameId int32
}

func NewClient(conn *WsConnection) *Client {
	c := &Client{
		cpack:       pack.CPackage{},
		spack:       pack.SPackage{},
		Conn:        conn,
		User:        entity.Users{},
		open:        true,
		pingTime:    time.Now().Unix(),
		grpcConnMap: make(map[string]*GrpcClientConn, 100),
	}

	//綁定函數
	c.Conn.wsConnect.SetPingHandler(func(appData string) error {
		byteNum := binary.Size(appData)
		if byteNum < 8 {
			return errors.New("非法的服務器")
		}
		temp := [2]byte{appData[0], appData[1]}
		num := bigendian.FromUint16(temp)
		if num != msg.CMD_HEART_BIT {
			return errors.New("非法的服務器")
		}

		fmt.Println("pong.............")
		sTime := int32(time.Now().Unix())
		heartToC := &gwpb.MHeartBitToc{Time: &sTime}
		data, _ := proto.Marshal(heartToC)

		c.Conn.wsConnect.WriteControl(websocket.PongMessage, data, time.Now().Add(writeWait))
		return nil
	})

	c.Conn.wsConnect.SetCloseHandler(func(code int, text string) error {
		var message []byte
		message = make([]byte, 4)

		tx := "異常關閉"
		if code == websocket.CloseGoingAway || code == websocket.CloseNormalClosure {
			tx = "客戶端正常關閉"
		}

		log.ZapLog.With(zap.Any("c", c)).Info("closed")
		if c != nil && c.sid != "" {
			log.ZapLog.With(zap.Any("userinfo", c.UserInfo)).Info(c.sid)
			userInfo, _ := service.NewUserService().GetUserInfo(c.sid)
			if userInfo != nil {
				userInfo.State = int64(0)
				service.NewUserService().UpdateUserInfo(c.UserInfo)
			}

			for _, grpcConn := range c.grpcConnMap {
				transId := fmt.Sprintf("%s_%xn", c.sid, grpcConn.url)

				roomId := int32(userInfo.RoomId)
				deskId := int32(userInfo.DeskId)
				tog := gwpb.MGame_1LeaveDeskTog{RoomId: &roomId, DeskId: &deskId}
				closeData, _ := proto.Marshal(&tog)

				grpcConn.FowardToBackServer(transId, uint32(msg.CMD_CLOSE), closeData)
			}
		}

		binary.BigEndian.PutUint32(message, uint32(code))
		message = append(message, []byte(tx)...)
		fmt.Printf("close.....%+d,%+v,%s", code, message, tx)
		c.Conn.wsConnect.WriteControl(websocket.CloseMessage, message, time.Now().Add(writeWait))

		return nil
	})

	return c
}

// Exec 客户端接收到数据 处理
func (c *Client) Exec() {
	var (
		err error
		msg pack.WsMessage
	)

	// 启动写协程
	go c.Conn.writeLoop()
	// 启动读协程
	go c.Conn.readLoop()

	c.handleSign()

	for {
		// 如果client已经关闭
		if c.open == false {
			break
		}

		// 如果websocket关闭
		if c.Conn.isClosed {
			c.Close()
			break
		}

		//接收数据  系统bug， 读取chan ,当没有<- 值的时候，会返回空长度
		msg, err = c.Conn.ReadMessage()
		if err != nil {
			log.ZapLog.With(zap.Any("error", err)).Info("收到错误")
			c.Close()
			break
		}

		//解包
		err = c.cpack.UnPkgBgData(msg)
		if err != nil {
			log.ZapLog.With(zap.Any("error", err)).Info("收到消息错误")
			continue
		}

		if c.cpack.ProtoNum != protoNum.CMD_HEART_BIT {
			info := fmt.Sprintf("收到cocos:%s,协议号=%+v,protobuf=%+v", c.ip, c.cpack.ProtoNum, c.cpack.ProtoData)
			log.ZapLog.Info(info)
		}

		pn := int(c.cpack.ProtoNum)
		log.ZapLog.Info(fmt.Sprintf("收到消息：=%+v", pn))

		if pn == int(protoNum.CMD_HEART_BIT) {
			c.sendPong()
		} else if pn == int(protoNum.CMD_REG) {
			reg := gwpb.MRegTos{}

			pbData := c.cpack.ProtoData
			if err := proto.Unmarshal(pbData, &reg); err != nil {
				log.ZapLog.With(zap.Any("err", err)).Info("replyReg")
				c.sendError(err)
				continue
			}

			if err = c.replyReg(reg); err != nil {
				c.sendError(err)
			}
		} else if pn == int(protoNum.CMD_LOGIN) {
			lts := gwpb.MLoginTos{}
			pd := c.cpack.ProtoData
			if err := proto.Unmarshal(pd, &lts); err != nil {
				log.ZapLog.With(zap.Any("err", err)).Info("SendLoginResult")
				continue
			}

			if err = c.sendLogin(lts); err != nil {
				c.sendError(err)
			}

			if err = c.sendGameConfig(); err != nil {
				c.sendError(err)
			}
			CManager.Register <- c
		} else {
			if c.sid == "" {
				c.sendError(errors.New("请先登录"))
				return
			}

			if err := c.proxyToGrpc(); err != nil {
				log.ZapLog.With(zap.Any("error", err)).Error("grpc错误")
				continue
			}
		}

		// 同时检测客户端ping 的超时是否
		n := time.Now().Unix()
		if n-c.pingTime > pongWaitSec {
			c.Conn.wsConnect.CloseHandler()(websocket.CloseNormalClosure, "服务器主动关闭")

			log.ZapLog.With(zap.Any("pingTime", c.pingTime), zap.Any("now", n)).Info("客户端连接超时，将退出!")
			c.Close()
			break
		}
	}

}

func (c *Client) sendPong() {
	sTime := int32(time.Now().Unix())
	heartToC := &gwpb.MHeartBitToc{Time: &sTime}
	data, _ := proto.Marshal(heartToC)

	d := c.spack.PkgSPackage(protoNum.CMD_HEART_BIT, data)

	c.Conn.WriteMessage(d)
	c.pingTime = int64(sTime)
}

func (c *Client) handleSign() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGPIPE, syscall.SIGHUP)
	go func() {
		for s := range sigs {
			switch s {
			//避免网络断开系统错误,进程挂掉
			case syscall.SIGPIPE, syscall.SIGHUP:
			default:
				log.ZapLog.With(zap.Any("Signal", s)).Info("other signal")
			}
		}
	}()
	return
}

func (c *Client) replyReg(l gwpb.MRegTos) error {
	userService := service.NewUserService()
	userInfoService := service.NewUserInfoService()
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(l.Pass))
	pass := hex.EncodeToString(md5Ctx.Sum(nil))

	if l.Pass != l.RePass {
		return errors.New("两次输入的密码不相等。")
	}

	if len(l.Pass) < 8 {
		return errors.New("密码长度必须等于大于8位。")
	}

	if l.Account == "" && len(l.Account) < 6 {
		return errors.New("账号不能够为空,必须6个字符串")
	}

	clientIp := c.Conn.wsConnect.RemoteAddr().String()
	user, err := userService.CreateAccount(l.Account, string(pass), clientIp)
	if err != nil {
		return err
	}

	_, err = userInfoService.CreateUserInfo(user.SId)
	if err != nil {
		return err
	}

	result := int32(1)
	lf := gwpb.MRegToc{
		Result: result,
	}
	data, _ := proto.Marshal(&lf)
	d := c.spack.PkgSPackage(protoNum.CMD_REG, data)

	log.ZapLog.With(zap.Any("protoNum.CMD_REG", protoNum.CMD_REG), zap.Any("data:", data)).Info("replyReg")
	err = c.Conn.WriteMessage(d)
	return err
}

func (c *Client) sendLogin(l gwpb.MLoginTos) error {
	redisClient := redislib.GetClient()
	var (
		u           entity.Users
		uf          entity.UserInfo
		err         error
		result      = int32(0) // 1 为失败 0 成功
		userService = service.NewUserService()
		token       string
	)

	pd := c.cpack.ProtoData
	if err := proto.Unmarshal(pd, &l); err != nil {
		log.ZapLog.With(zap.Any("err", err)).Info("SendLoginResult")
		return errors.New("proto3解码错误")
	}

	//查询数据库 是否有效的用户
	if u, err = userService.GetUserByPwd(l.Account, "", l.Pass); err == nil {
		result = int32(0)
		tokenService := service.TokenService{}
		token, err = tokenService.Sign(u, 3600*24)
		if err == nil {
			u.Token = token
			userService.UpdateUser(u)
		}

		log.ZapLog.With(zap.Namespace("ns"), zap.Any("Token", token), zap.Any("err", err)).Warn("token ")
	} else {
		log.ZapLog.With(zap.Any("err", err)).Info("SendLoginResult 登录失败")
		result = int32(1)
	}

	// 发送是否登录成功信息
	lf := gwpb.MLoginToc{
		Result: result,
		Token:  &token,
	}
	data, _ := proto.Marshal(&lf)
	d := c.spack.PkgSPackage(protoNum.CMD_LOGIN, data)
	c.Conn.WriteMessage(d)

	if err != nil {
		return nil
	}

	//设置用户redis
	userService.SetUserToRds(u, redisClient)

	// 如果登录成功 发送用户信息
	mu := gwpb.MUserToc{}
	copier.Copy(&mu, &u)
	//mu := (*gwpb.MUser)(unsafe.Pointer(&u))
	data, _ = proto.Marshal(&mu)
	d = c.spack.PkgSPackage(protoNum.CMD_USER, data)
	err = c.Conn.WriteMessage(d)

	// 如果登录成功 发送玩家信息
	uf.State = 1
	muf := gwpb.MUserInfoToc{}
	copier.Copy(&muf, &uf)
	data, _ = proto.Marshal(&muf)
	d = c.spack.PkgSPackage(protoNum.CMD_USER_INFO, data)
	err = c.Conn.WriteMessage(d)

	// 更新为在线
	service.NewUserService().UpdateUserInfo(uf)

	c.UserInfo = uf
	c.User = u
	c.sid = c.User.SId
	return err
}

func (c *Client) sendGameConfig() (err error) {
	pgcs := make([]*gwpb.PGameConfig, 0)
	grs := make([]*gwpb.PRoom, 0)

	games := make([]entity.Games, 0)
	err = mysql.S1().Table(entity.TABLE_GAME).Select("id,is_delete").Find(&games)
	if err != nil {
		log.ZapLog.With(zap.Namespace("database"), zap.Any("err", err)).Error("数据库查询错误")
	}

	for _, g := range games {
		pgc := &gwpb.PGameConfig{}

		rs := make([]entity.Room, 0)
		err = mysql.S1().Table(entity.TABLE_ROOM).Select("id,score,status").Where("game_id=?", g.ID).Find(&rs)
		for _, r := range rs {
			lr := gwpb.PRoom{}
			copier.Copy(&lr, &r)
			grs = append(grs, &lr)
		}

		pgc.RoomInfo = grs
		pgc.GameId = int32(g.ID)
		pgc.State = int32(g.IsDelete)
		pgcs = append(pgcs, pgc)
	}

	mgct := gwpb.MGameConfigToc{
		Game: pgcs,
	}

	log.ZapLog.With(zap.Any("mgct", mgct)).Info("gameConfig")
	data, _ := proto.Marshal(&mgct)
	d := c.spack.PkgSPackage(protoNum.CMD_GAME_CONFIG, data)
	err = c.Conn.WriteMessage(d)

	return err
}

func (c *Client) sendError(err error) {
	sTime := int32(time.Now().Unix())
	code := int32(0)
	errToC := &gwpb.MServerErrorToc{
		Code: &code,
		Text: err.Error(),
	}
	data, _ := proto.Marshal(errToC)

	d := c.spack.PkgSPackage(protoNum.CMD_ERROR, data)
	c.Conn.WriteMessage(d)
	c.pingTime = int64(sTime)
}

//todo 转发给各个服务器
func (c *Client) proxyToGrpc() (err error) {
	pn := int(c.cpack.ProtoNum)

	log.ZapLog.With(zap.Any("pn", pn)).Info("proxyToGrpc")
	backUrl := getBackUrl(pn)
	log.ZapLog.With(zap.Any("backUrl", backUrl)).Info("proxyToGrpc")
	if backUrl == "" {
		return errors.New("没有找不到后端服务器，请检查代理服务器配置")
	}

	// todo 这里需要判断协议
	if strings.Contains(backUrl, "grpc") {
		b := []byte(backUrl)
		url := string(b[7:])
		// todo 必须做到多路复用， 多路复用使用连接池。否则性能太差了
		// todo 必须妥善处理grpc断线重连 重拨， 在grpc stream 收发数据的之前，需要判断grpc服务器的状态。 需要做一个grpc dial 管理的类
		// todo 为确定服务器状态,当服务器关闭的时候,需要推送流到gateway, 作为服务器马上关闭的信号,gateway收到做处理
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(backUrl))
		transId := fmt.Sprintf("%s_%xn", c.sid, string(md5Ctx.Sum([]byte(""))))
		gpls, ok := PoolsCollect[url]

		log.ZapLog.With(zap.Any("transId", transId)).Info("proxyToGrpc.....")
		if ok {
			c.currGrpcConn, err = gpls.Get(transId)
			if err != nil {
				c.sendError(err)
			}
			c.currGrpcConn.AddClient(transId)

			Lock.Lock()
			Clients[transId] = *c
			c.currGrpcConn.ClientsKeys = append(c.currGrpcConn.ClientsKeys, transId)
			Lock.Unlock()

			log.ZapLog.With(zap.Int32("LinkClientNum", c.currGrpcConn.GetLinkClientNum())).Info("grpc.....")
		} else {
			log.ZapLog.With(zap.Any("transId", "没有初始化对应grpc client 连接池")).Info("proxyToGrpc.....")
			return errors.New("没有初始化对应grpc client 连接池")
		}

		cp := c.cpack
		ptn := uint32(cp.ProtoNum)

		// 转发给各个服务器
		//secret := []byte{cp.Secret[0], cp.Secret[1]}
		//serialNum := []byte{cp.RandNum[0], cp.RandNum[1], cp.RandNum[2], cp.RandNum[3]}
		//secret := []byte{'0', '0'}
		//serialNum := []byte{'0', '0', '0', '0'}
		data := cp.ProtoData
		if err := c.currGrpcConn.FowardToBackServer(transId, ptn, data); err != nil {
			log.ZapLog.With(zap.Any("err", err)).Info("FowardToBackServer")

			c.sendError(err)
		}
	} else {
		return errors.New("不支持grpc外的协议")
	}

	return nil
}

func (c *Client) Open() {
	c.rwlock.Lock()
	defer c.rwlock.Unlock()
	c.open = true
}

// Close 释放资源
func (c *Client) Close() {
	c.rwlock.RLock()
	defer c.rwlock.RUnlock()

	c.CManager.Unregister <- c
	c.open = false
	c.Conn.Close()
	c.currGrpcConn.Close()
}

func getBackUrl(ProtoNum int) string {
	u := ""
	for _, proxy := range viper.PsCfg {
		hsp := proxy.StartProtocal
		hep := proxy.EndProtocal
		for _, url := range proxy.Addr {
			if ProtoNum >= hsp && ProtoNum <= hep {
				u = url
				goto RET
			}
		}
	}
RET:
	return u
}

func (c *Client) sendLoginOldVersion(l gwpb.MLoginTos) error {
	var (
		u      *entity.Users
		uf     *entity.UserInfo
		result = int32(0) // 1 为失败 0 成功
		err    error
	)
	pd := c.cpack.ProtoData
	if err := proto.Unmarshal(pd, &l); err != nil {
		log.ZapLog.With(zap.Any("err", err)).Info("SendLoginResult")
		return errors.New("proto3解码错误")
	}

	sid := strconv.FormatInt(11111222, 10)
	//sid := l.Account
	if u, uf, err = GetUserInfo(sid); err != nil || uf == nil || u == nil {
		// 登录失败
		result = int32(1)
		lf := gwpb.MLoginToc{
			Result: result,
			Token:  nil,
		}
		data, _ := proto.Marshal(&lf)
		d := c.spack.PkgSPackage(protoNum.CMD_LOGIN, data)

		log.ZapLog.With(zap.Any("err", err)).Error("GetUserInfo")
		c.Conn.WriteMessage(d)
		return err
	}

	// 登录成功信息
	loginProto := gwpb.MLoginToc{
		Result: result,
		Token:  &l.Pass,
	}
	data, _ := proto.Marshal(&loginProto)
	d := c.spack.PkgSPackage(protoNum.CMD_LOGIN, data)
	err = c.Conn.WriteMessage(d)

	// 如果登录成功 发送用户信息
	mu := gwpb.MUserToc{}
	copier.Copy(&mu, &u)
	//mu := (*gwpb.MUser)(unsafe.Pointer(&u))
	data, _ = proto.Marshal(&mu)
	d = c.spack.PkgSPackage(protoNum.CMD_USER, data)
	err = c.Conn.WriteMessage(d)

	// 如果登录成功 发送玩家信息
	uf.State = 1
	muf := gwpb.MUserInfoToc{}
	copier.Copy(&muf, &uf)
	data, _ = proto.Marshal(&muf)
	d = c.spack.PkgSPackage(protoNum.CMD_USER_INFO, data)
	err = c.Conn.WriteMessage(d)

	// 更新为在线
	service.NewUserService().UpdateUserInfo(*uf)

	c.UserInfo = *uf
	c.User = *u
	c.sid = c.User.SId

	return err
}
