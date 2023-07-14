package bj_server

import (
	"time"
)

type BroadcastConfig struct {
	Id                     uint      `xorm:"not null pk autoincr UNSIGNED INT"`
	Name                   string    `xorm:"not null comment('广播名称') VARCHAR(255)"`
	Type                   int       `xorm:"not null comment('广播类型1：平台2：游戏') TINYINT"`
	GameId                 int       `xorm:"not null default 0 comment('游戏') TINYINT"`
	Content                string    `xorm:"not null comment('广播内容') TEXT"`
	PlayerTriggerGold      int64     `xorm:"not null comment('触发金额') BIGINT"`
	RobotTriggerMinGold    int64     `xorm:"not null comment('机器人最小触发金额') BIGINT"`
	RobotTriggerMaxGold    int64     `xorm:"not null comment('机器人最大触发金额') BIGINT"`
	FakeBroadcastStartTime int       `xorm:"not null comment('假广播开始时间') INT"`
	FakeBroadcastEndTime   int       `xorm:"not null comment('假广播结束时间') INT"`
	CreatedAt              time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UpdatedAt              time.Time `xorm:"not null comment('更新时间') DATETIME"`
}

func (m *BroadcastConfig) TableName() string {
	return "broadcast_config"
}
