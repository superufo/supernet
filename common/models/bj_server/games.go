package bj_server

import (
	"time"
)

type Games struct {
	Id          uint      `xorm:"not null pk UNSIGNED SMALLINT"`
	Name        string    `xorm:"not null default '' comment('游戏名称') VARCHAR(60)"`
	Icon        string    `xorm:"not null default '' comment('游戏图标') VARCHAR(200)"`
	Code        string    `xorm:"not null default '' comment('游戏编码') VARCHAR(30)"`
	LocaleNames string    `xorm:"not null comment('游戏名称语言配置') JSON"`
	IsDelete    uint      `xorm:"not null default 0 comment('是否删除') UNSIGNED TINYINT"`
	CreatedAt   time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *Games) TableName() string {
	return "games"
}
