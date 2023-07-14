package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"

	"go.uber.org/zap"

	"github.com/supernet/common/utils"
	"github.com/supernet/gateway/internal/model/entity"

	"github.com/supernet/gateway/pkg/log"
	"github.com/supernet/gateway/pkg/mysql"
	"github.com/supernet/gateway/pkg/redislib"
)

// UserService  获取 user user_info 表格中信息
type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us UserService) UpdateUser(user entity.Users) error {
	_, err := mysql.M().Table(entity.TABLE_USERS).Update(&user)

	return err
}

func (us UserService) UpdateUserInfo(user entity.UserInfo) error {
	_, err := mysql.M().Table(entity.TABLE_USER_INFO).Update(user)

	return err
}

func (us UserService) GetUserByPwd(account string, mac string, pwd string) (user entity.Users, err error) {
	redislib.Sclient()
	redisClient := redislib.GetClient()
	defer redisClient.Close()

	ukey := fmt.Sprintf("user:%s", user.Account)
	if err := redisClient.HGetAll(context.Background(), ukey).Scan(&user); err != nil {
		log.ZapLog.With(zap.Namespace("redis"), zap.Any("err", err), zap.Any("ukey", ukey)).Error("redis操作错误")
	}

	isOK, err := mysql.S1().Table(entity.TABLE_USERS).Select("*").Where("(account=? or mac =?) and password=?", account, mac, pwd).Get(&user)
	if err != nil {
		sql, _ := mysql.S1().Table(entity.TABLE_SERVER_LIST).LastSQL()
		log.ZapLog.With(zap.Namespace("database"), zap.Any("err", err), zap.Any("sql", sql)).Error("数据库查询错误")
	}

	//todo 判断是否为游客
	if isOK == false {
		user, _ = us.GeneralGustAccount(mac, pwd)
	}

	return user, err
}

func (us UserService) CreateAccount(account string, pwd string, ip string) (user entity.Users, err error) {
	log.ZapLog.With(zap.Any("account：", account), zap.Any("pwd：", pwd)).Info("插入数据")

	var affected int64
	user = entity.Users{
		SId:          utils.GenRandString(16),
		Account:      account,
		Password:     pwd, //string(fmt.Sprintf("%x", md5.Sum([]byte{1, 2, 3, 4, 5, 6}))),
		RegisterTime: time.Now().Unix(),
		RegisterIp:   ip,
	}

	if affected, err = mysql.M().Table(entity.TABLE_USERS).Insert(user); err == nil {
		if affected < 1 {
			log.ZapLog.With(zap.Any("affected", affected)).Warn("数据库插入数据为空")
		}
	} else {
		log.ZapLog.With(zap.Any("err", err)).Warn("数据库插入错误")
	}

	return user, err
}

func (us UserService) GeneralGustAccount(mac string, pwd string) (user entity.Users, err error) {
	rand.Seed(time.Now().Unix())

	user = entity.Users{
		SId:          utils.GenRandString(16),
		ID:           rand.Intn(1000),
		Account:      fmt.Sprintf("account_%s", utils.GenRandString(16)),
		Name:         fmt.Sprintf("guest_%s", utils.GenRandString(16)),
		Token:        "",
		Platform:     "",
		Sex:          int8(rand.Intn(1)),
		Mac:          mac, //utils.GenRandString(32),
		Nickname:     fmt.Sprintf("guest_%s", utils.GenRandString(16)),
		CCode:        "",
		Phone:        "",
		RegisterTime: time.Now().Unix(),
		Password:     pwd, //string(fmt.Sprintf("%x", md5.Sum([]byte{1, 2, 3, 4, 5, 6}))),
		Agent:        "",
		Status:       int8(0),
		RegisterIp:   "",
		FatherId:     "",
	}

	if affected, err := mysql.M().Table(entity.TABLE_USERS).Insert(user); err == nil {
		if affected < 1 {
			log.ZapLog.With(zap.Any("affected", affected)).Error("数据库插入数据为空")
		}
	} else {
		log.ZapLog.With(zap.Any("err", err)).Error("数据库插入错误")
	}

	return user, err
}

func (us UserService) SetUserToRds(user entity.Users, c *redis.Client) (bool, error) {
	key := fmt.Sprintf("user:%s", user.SId)
	ctx := context.Background()
	res := c.HMSet(
		ctx, key,
		"s_id", user.SId,
		"id", user.ID,
		"account", user.Account,
		"name", user.Name,
		"token", user.Token,
		"platform", user.Platform,
		"sex", user.Sex,
		"mac", user.Mac,
		"nickname", user.Nickname,
		"c_code", user.CCode,
		"phone", user.Phone,
		"register_time", user.RegisterTime,
		"password", user.Password,
		"agent", user.Agent,
		"status", user.Status,
		"register_ip", user.RegisterIp,
		"father_id", user.FatherId,
	)

	return res.Result()
}

func (us UserService) GetUserByToken(token string, account string) (user entity.Users, err error) {
	var has bool
	if token == "" {
		err = errors.New("token为空")
		goto RET
	}

	if account == "" {
		err = errors.New("account为空")
		goto RET
	}

	has, err = mysql.S1().Table(entity.TABLE_USERS).Where("token=?  and account=?", token, account).Get(&user)
	if err != nil {
		log.ZapLog.With(zap.Namespace("database"), zap.Any("err", err)).Error("数据库查询错误")
	}

	if !has {
		err = errors.New("用户不存在")
		goto RET
	}

RET:
	return user, err
}

func (us UserService) GetUserInfo(sid string) (*entity.UserInfo, error) {
	redis := redislib.GetClient()
	ctx := context.Background()
	key := fmt.Sprintf("userInfo:%s", sid)
	info := entity.UserInfo{}
	//if err := redis.HGetAll(ctx, key).Scan(&info); err != nil {
	//	log.ZapLog.Info("redis error", zap.Any("err", err), zap.Any("key", key))
	//}
	if info.SId == "" {
		ok, err := mysql.S1().Table(entity.TABLE_USER_INFO).Where("s_id=?", sid).Get(&info)
		if err != nil {
			sql, _ := mysql.S1().Table(entity.TABLE_USER_INFO).LastSQL()
			log.ZapLog.Error("数据库查询错误", zap.Any("database err", err), zap.Any("sql", sql))
			return nil, err
		}
		if ok {
			redis.HMSet(
				ctx, key,
				"s_id", info.SId,
				"login_time", info.LoginTime,
				"offline_time", info.OfflineTime,
				"gold", info.Gold,
				"diamonds", info.Diamonds,
				"state", info.State,
				"login_ip", info.LoginIp,
				"login_s_flag", info.LoginSFlag,
				"ctrl_status", info.CtrlStatus,
				"game_id", info.GameId,
				"room_id", info.RoomId,
				"desk_id", info.DeskId,
				"ctrl_value", info.CtrlValue,
				"p_stock", info.PStock,
				"recent_play_time", info.RecentPlayTime,
				"total_recharge", info.TotalRecharge,
				"total_cash", info.TotalCash,
				"gm_award_1", info.GmAward1,
				"gm_award_2", info.GmAward2,
				"recent_play_per_round_sid", info.RecentPlayPerRoundSid,
				"ctrl_data", info.CtrlData,
				"ctrl_probability", info.CtrlProbability,
				"ctrl_scales", info.CtrlScales,
				"platform", info.Platform,
				"agent", info.Agent,
			)
		} else {
			return nil, fmt.Errorf("GetUserInfo can not find user:%v", sid)
		}
	}
	log.ZapLog.Info("GetUserInfo", zap.Any("key", key), zap.Any("info", info))

	return &info, nil
}
