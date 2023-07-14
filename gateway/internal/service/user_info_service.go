package service

import (
	"fmt"
	"go.uber.org/zap"

	"context"
	"github.com/supernet/gateway/internal/model/entity"

	"github.com/supernet/gateway/pkg/log"
	"github.com/supernet/gateway/pkg/mysql"
	"github.com/supernet/gateway/pkg/redislib"
)

// UserService  获取 user user_info 表格中信息
type UserInfoService struct {
}

func NewUserInfoService() *UserInfoService {
	return &UserInfoService{}
}

func (us UserInfoService) CreateUserInfo(sid string) (*entity.UserInfo, error) {
	info := entity.UserInfo{SId: sid}

	var (
		affected int64
		err      error
	)

	if affected, err = mysql.M().Table(entity.TABLE_USER_INFO).Insert(info); err == nil {
		if affected < 1 {
			log.ZapLog.With(zap.Any("affected", affected)).Warn("数据库插入数据为空")
		}
	} else {
		log.ZapLog.With(zap.Any("err", err)).Warn("数据库插入错误")
	}

	return &info, nil
}

func (us UserInfoService) GetUserInfo(sid string) (*entity.UserInfo, error) {
	redis := redislib.GetClient()
	ctx := context.Background()
	key := fmt.Sprintf("userInfo:%s", sid)
	info := entity.UserInfo{}

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
