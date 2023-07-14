package bj_server

import (
	"time"
)

type RobotConfig struct {
	Id        uint      `xorm:"not null pk autoincr UNSIGNED SMALLINT"`
	GameId    uint      `xorm:"not null default 0 comment('游戏id') UNSIGNED SMALLINT"`
	Conf      string    `xorm:"not null default '' comment('机器人配置数据') VARCHAR(3000)"`
	CreatedAt time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	UpdatedAt time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP"`
}

func (m *RobotConfig) TableName() string {
	return "robot_config"
}
