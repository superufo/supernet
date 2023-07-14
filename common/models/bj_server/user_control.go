package bj_server

import (
	"time"
)

type UserControl struct {
	Id                 uint      `xorm:"not null pk autoincr UNSIGNED INT"`
	SId                string    `xorm:"not null default '' comment('用户服务器id(唯一)') unique(idx_s_id_del) VARCHAR(32)"`
	IntegerUserId      uint      `xorm:"not null comment('用户整型id(唯一)') index(idx_id_del) UNSIGNED INT"`
	ControlAmount      int64     `xorm:"not null default 0 comment('控制金额') BIGINT"`
	ControlAmountRatio uint      `xorm:"not null default 0 comment('控制金额浮动比例') UNSIGNED TINYINT"`
	ControlProbability uint      `xorm:"not null default 0 comment('控制概率') UNSIGNED TINYINT"`
	Platform           uint      `xorm:"not null default 0 comment('用户渠道id') UNSIGNED INT"`
	Agent              string    `xorm:"not null default '' comment('用户代理标识') VARCHAR(128)"`
	IsDelete           uint      `xorm:"not null default 0 comment('是否删除') index(idx_id_del) unique(idx_s_id_del) UNSIGNED TINYINT"`
	CreatedAt          time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdatedAt          time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *UserControl) TableName() string {
	return "user_control"
}
