package bj_log

import (
	"time"
)

type UserGameIeDayStat struct {
	Id                  uint      `xorm:"not null pk autoincr UNSIGNED INT"`
	Day                 time.Time `xorm:"not null comment('日期') index(idx_day_cid_iid_gid) unique(uk_day_sid_gid) DATE"`
	SId                 string    `xorm:"not null default '' comment('用户服务器id') unique(uk_day_sid_gid) VARCHAR(32)"`
	IntegerUserId       uint      `xorm:"not null default 0 comment('用户整型id') index(idx_day_cid_iid_gid) UNSIGNED INT"`
	ChannelId           uint      `xorm:"not null default 0 comment('渠道id') index(idx_day_cid_iid_gid) unique(uk_day_sid_gid) UNSIGNED INT"`
	GameId              uint      `xorm:"not null default 0 comment('游戏id') index(idx_day_cid_iid_gid) unique(uk_day_sid_gid) UNSIGNED SMALLINT"`
	BetsGold            uint64    `xorm:"not null default 0 comment('下注金额') UNSIGNED BIGINT"`
	SettleGold          uint64    `xorm:"not null default 0 comment('结算金额') UNSIGNED BIGINT"`
	BalanceGold         int64     `xorm:"not null default 0 comment('收支金额') BIGINT"`
	WinGold             int64     `xorm:"not null default 0 comment('赢钱金额') BIGINT"`
	LoseGold            int64     `xorm:"not null default 0 comment('输钱金额') BIGINT"`
	PlayerServiceCharge int64     `xorm:"not null default 0 comment('玩家服务费') BIGINT"`
	CompanyLoseGold     int64     `xorm:"not null default 0 comment('公司(平台)损失金币(玩家输钱时触发)') BIGINT"`
	ChannelLoseGold     int64     `xorm:"not null default 0 comment('渠道损失金币(玩家输钱时触发)') BIGINT"`
}

func (m *UserGameIeDayStat) TableName() string {
	return "user_game_ie_day_stat"
}
