package net

import (
	"context"
	"errors"
	"fmt"
	"github.com/supernet/gateway/internal/net/gstream/pb"
	"github.com/supernet/gateway/pkg/log"
	"github.com/supernet/gateway/pkg/pack"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

var (
	retryPolicy = `{
		"methodConfig": [{
		  "name": [{"service": "backurl"}],
		  "waitForReady": true,
		  "retryPolicy": {
			  "MaxAttempts": 5,
              "hedgingDelay": "0.05s",  
			  "InitialBackoff": ".01s",
			  "MaxBackoff": ".01s",
			  "BackoffMultiplier": 1.0,
			  "RetryableStatusCodes": [ "UNAVAILABLE" ]
		  }
		}]}`

	Lock    = sync.Mutex{}
	Clients = make(map[string]Client, 100) // 全局的client数量
)

type GrpcClientConn struct {
	GrpcRecvData chan *pb.StreamResponseData
	GrpcSendData chan *pb.StreamRequestData
	// 死星管道
	GrpcDeadData    chan *pb.StreamRequestData
	clientConn      *grpc.ClientConn
	rpc             pb.ForwardMsgClient
	clientStream    pb.ForwardMsg_PPStreamClient
	linkClientNum   int32
	url             string
	lock            sync.Mutex
	ctx             context.Context
	ctxCancle       context.CancelFunc
	ClientsKeys     []string
	isDisConnection bool
}

func NewGrpcClientConn(url string) (*GrpcClientConn, error) {
	var (
		clientStream pb.ForwardMsg_PPStreamClient
		err          error
	)

	diaOpt := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(2000*1024*1024),
		grpc.MaxCallSendMsgSize(2000*1024*1024))
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		//grpc.WithKeepaliveParams(kacp),
		grpc.WithDefaultServiceConfig(retryPolicy),
		diaOpt)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("rpc连接错误:%s" + err.Error())) // errors.New("rpc连接错误:"+err )
	}
	rpc := pb.NewForwardMsgClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	if conn != nil {
		if clientStream, err = rpc.PPStream(ctx); err != nil {
			log.ZapLog.With(zap.Any("err", err)).Info("Run")
			return nil, err
		}
	}

	gcn := &GrpcClientConn{
		GrpcRecvData:    make(chan *pb.StreamResponseData, 50),
		GrpcSendData:    make(chan *pb.StreamRequestData, 50),
		linkClientNum:   int32(0),
		clientConn:      conn,
		rpc:             rpc,
		clientStream:    clientStream,
		url:             url,
		ctx:             ctx,
		ctxCancle:       cancel,
		lock:            sync.Mutex{},
		ClientsKeys:     make([]string, 0),
		isDisConnection: false,
	}

	//log.ZapLog.With(zap.Any("conn.GetState", conn.GetState())).Info("NewGrpcClientConn")
	return gcn, nil
}

func (gc *GrpcClientConn) GetLinkClientNum() int32 {
	return gc.linkClientNum
}

func (gc *GrpcClientConn) SubLinkClientNum() int32 {
	return atomic.AddInt32(&gc.linkClientNum, -1)
}

func (gc *GrpcClientConn) revStream() {
	for {
		msg, err := gc.clientStream.Recv()

		log.ZapLog.With(zap.Any("state", gc.clientConn.GetState())).Info("ztai")
		if err == io.EOF {
			return
		}

		if err != nil {
			log.ZapLog.With(zap.Any("error", err)).Error("Recv")
			gc.Close()
			return
		}

		if msg.GetMsg() == 0 {
			continue
		}

		log.ZapLog.With(zap.Any("msg", msg.GetMsg()), zap.Any("data", msg.GetData())).Info("rev backend")
		gc.GrpcRecvData <- msg
	}
}

// log.ZapLog.With(zap.Any("err", err)).Error("Send调用grpc方法错误")
func (gc *GrpcClientConn) sendStream() {
	var sd *pb.StreamRequestData

	for {
		select {
		case sd = <-gc.GrpcSendData:
			err := gc.clientStream.Send(sd)

			if err == io.EOF || err != nil {
				// 发送失败放入死信管道
				gc.GrpcDeadData <- sd
				gc.Close()
				log.ZapLog.With(zap.Any("error", err)).Error("Send")
				return
			}

			log.ZapLog.With(zap.Any("msg", sd.GetMsg()), zap.Any("data", sd.GetData())).Info("to backend")
			//zap.Any("secret", sd.GetSecret()), zap.Any("serialNum", sd.GetSerialNum())).Info("to backend")
		}
	}
}

func (gc *GrpcClientConn) proxyToClient() {
	for {
		select {
		case grc := <-gc.GrpcRecvData:
			// currGrpcConn 对应多个 clientId, 只做等于自己的clientId的转发
			if grc == nil || grc.Msg == 0 {
				continue
			}

			log.ZapLog.With(zap.Any("grc.Msg", grc.Msg), zap.Any("grc.ClientId", grc.ClientId), zap.Any("grc.BAllUser", grc.BAllUser),
				zap.Any("grc.Uids", grc.Uids)).Info("后端服务器转发给网关消息")

			// 通过grc.ClientId 来获取sid
			var sid string
			if grc.ClientId != "" {
				sid = strings.Split(grc.ClientId, "-")[0]
			}

			p := pack.NewSPackage()
			p.PkgSPackage(uint16(grc.Msg), grc.Data)
			if grc.BAllUser {
				for uid, client := range CManager.Clients {
					client.Conn.WriteMessage(p.PkgSPackage(uint16(grc.Msg), grc.Data))
					log.ZapLog.With(zap.Any("uid", uid), zap.Any("Msg", uint16(grc.Msg))).Info("广播reply client.....")
				}
			} else if len(grc.Uids) > 0 {
				for _, uid := range grc.Uids {
					if client, ok := CManager.Clients[uid]; ok {
						client.Conn.WriteMessage(p.PkgSPackage(uint16(grc.Msg), grc.Data))
						log.ZapLog.With(zap.Any("uid", uid), zap.Any("Msg", uint16(grc.Msg))).Info("组播 reply client.....")
					}
				}
			} else if sid != "" {
				if client, ok := Clients[grc.ClientId]; ok {
					client.Conn.WriteMessage(p.PkgSPackage(uint16(grc.Msg), grc.Data))
					log.ZapLog.With(zap.Any("grpc clientId", grc.ClientId), zap.Any("Msg", uint16(grc.Msg))).Info("reply client.....")
				} else {
					log.ZapLog.With(zap.Any("grpc clientId", grc.ClientId)).Error("grpc回复消息没有对应的client存在")
				}
			} else {
				log.ZapLog.With(zap.Any("grpc clientId", grc.ClientId), zap.Any("Msg", uint16(grc.Msg))).Info("没有用户信息 ")
			}
		}
	}
}

func (gc *GrpcClientConn) FowardToBackServer(clientId string, msg uint32, data []byte) (err error) {
	srd := pb.StreamRequestData{
		ClientId: clientId,
		Msg:      msg,
		//Secret:    secret,
		//SerialNum: serialNum,
		Data: data,
	}

	//gc.GrpcSendData <- &srd
	//info := fmt.Sprintf("转发给后端:协议号=%+v,加密字符=%+v,随机字符=%+v,protobuf=%+v", msg, Utils.ToHexString(secret), serialNum, data)
	//log.ZapLog.Info(info)
	if (gc.clientConn != nil) && !gc.disConnection() {
		gc.GrpcSendData <- &srd

		log.ZapLog.With(zap.Any("srd", &srd)).Info("FowardToBackServer")
	} else {
		err = errors.New("后端服务器已经关闭")
	}

	return err
}

func (gc *GrpcClientConn) Run() error {
	var err error

	// grpc 从服务器收发
	go gc.sendStream()
	go gc.revStream()

	//转发到client端
	go gc.proxyToClient()

	// syscall.SIGHUP 挂断
	// syscall.SIGPIPE 往断开的通道读或者写
	// todo  应该新启动一个管理类来做grpc的这些action
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGPIPE, syscall.SIGHUP)
	go func() {
		for s := range c {
			switch s {
			//避免网络断开系统错误,进程挂掉
			case syscall.SIGPIPE, syscall.SIGHUP:
				log.ZapLog.Info("收到信号SIGPIPE or  SIGHUP signal")
				gc.isDisConnection = true
				// 获取连接池
				if pools, has := PoolsCollect[gc.url]; has {
					for _, ck := range gc.ClientsKeys {
						log.ZapLog.With(zap.Any("client key", ck)).Info("释放cleint bind")
						pools.Return(ck)
					}
				} else {
					log.ZapLog.With().Error("连接池获取失败")
				}

				timer := time.NewTicker(time.Millisecond * 100)
				for {
					select {
					case <-timer.C:
						gc, err = NewGrpcClientConn(gc.url)
						if err == nil {
							log.ZapLog.With().Info("创建成功")
							gc.Run()
							break
						} else {
							log.ZapLog.With(zap.Error(err)).Error("创建失败")
						}
					}
				}
			default:
				log.ZapLog.With(zap.Any("Signal", s)).Info("other signal")
			}
		}
	}()

	return err
}

func (gc *GrpcClientConn) AddClient(clientKey string) {
	gc.ClientsKeys = append(gc.ClientsKeys, clientKey)
}

func (gc *GrpcClientConn) NeedClose() bool {
	gstate := gc.clientConn.GetState()
	if gstate == connectivity.TransientFailure || gstate == connectivity.Idle {
		return true
	}

	return false
}

func (gc *GrpcClientConn) disConnection() bool {
	gstate := gc.clientConn.GetState()
	if gstate == connectivity.Shutdown || gstate == connectivity.TransientFailure {
		return true
	}

	return false
}

func (gc *GrpcClientConn) Close() {
	if gc != nil && !gc.isDisConnection {
		gc.lock.Lock()
		gc.isDisConnection = true

		if _, ok := <-gc.GrpcRecvData; ok != true {
			close(gc.GrpcRecvData)
		}

		if _, ok := <-gc.GrpcSendData; ok != true {
			close(gc.GrpcSendData)
		}

		gc.clientConn.Close()
		gc.ctxCancle()
		gc.lock.Unlock()
	}
}
