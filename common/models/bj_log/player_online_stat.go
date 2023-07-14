package bj_log

import (
	"time"
)

type PlayerOnlineStat struct {
	Id          int       `xorm:"not null pk autoincr INT"`
	ChannelId   int       `xorm:"not null comment('渠道ID') INT"`
	Game1Online int       `xorm:"not null default 0 comment('龙虎在线人数') INT"`
	Game2Online int       `xorm:"not null default 0 comment('红黑在线人数') INT"`
	Game3Online int       `xorm:"not null default 0 comment('百家乐在线人数') INT"`
	Game4Online int       `xorm:"not null default 0 comment('鱼虾蟹在线人数') INT"`
	StatTime    time.Time `xorm:"not null comment('统计时间') DATETIME"`
}

func (m *PlayerOnlineStat) TableName() string {
	return "player_online_stat"
}
