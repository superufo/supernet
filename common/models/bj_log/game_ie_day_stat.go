package bj_log

import (
	"time"
)

type GameIeDayStat struct {
	Id          uint      `xorm:"not null pk autoincr UNSIGNED INT"`
	Day         time.Time `xorm:"not null comment('日期') unique(uk_day_gid_cid) DATE"`
	GameId      uint      `xorm:"not null default 0 comment('游戏id') unique(uk_day_gid_cid) UNSIGNED SMALLINT"`
	ChannelId   uint      `xorm:"not null default 0 comment('渠道id') unique(uk_day_gid_cid) UNSIGNED INT"`
	BetsGold    uint64    `xorm:"not null default 0 comment('下注金额') UNSIGNED BIGINT"`
	SettleGold  uint64    `xorm:"not null default 0 comment('结算金额') UNSIGNED BIGINT"`
	BalanceGold int64     `xorm:"not null default 0 comment('收支金额') BIGINT"`
	BetsUsers   uint      `xorm:"not null default 0 comment('下注人数') UNSIGNED INT"`
	WinGold     int64     `xorm:"not null default 0 comment('赢钱金额') BIGINT"`
	LoseGold    int64     `xorm:"not null default 0 comment('输钱金额') BIGINT"`
}

func (m *GameIeDayStat) TableName() string {
	return "game_ie_day_stat"
}
