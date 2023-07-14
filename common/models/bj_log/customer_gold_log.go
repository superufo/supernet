package bj_log

import (
	"time"
)

type CustomerGoldLog struct {
	Id            uint      `xorm:"not null pk autoincr UNSIGNED INT"`
	SId           string    `xorm:"not null default '' comment('玩家服务器id(唯一)') index VARCHAR(32)"`
	IntegerUserId uint      `xorm:"not null default 0 comment('玩家整型id(唯一)') index UNSIGNED INT"`
	Type          uint      `xorm:"not null default 0 comment('操作类型 1.存入 2.扣除') UNSIGNED TINYINT"`
	Reason        uint      `xorm:"not null default 0 comment('操作原因') UNSIGNED TINYINT"`
	BeforeGold    int64     `xorm:"not null default 0 comment('操作前玩家金币数') BIGINT"`
	Gold          int64     `xorm:"not null default 0 comment('操作金币数') BIGINT"`
	AfterGold     int64     `xorm:"not null default 0 comment('操作后玩家金币数') BIGINT"`
	OperatorId    uint      `xorm:"not null default 0 comment('操作用户id') UNSIGNED INT"`
	Remark        string    `xorm:"not null default '' comment('操作备注') VARCHAR(255)"`
	CreatedAt     time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *CustomerGoldLog) TableName() string {
	return "customer_gold_log"
}
