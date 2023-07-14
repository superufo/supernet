package bj_agent

import (
	"time"
)

type Users struct {
	Id              uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT"`
	ChannelName     string    `xorm:"not null comment('渠道名称') unique VARCHAR(255)"`
	Name            string    `xorm:"not null comment('账号名称') unique VARCHAR(255)"`
	Avatar          string    `xorm:"not null comment('头像') VARCHAR(255)"`
	Phone           int64     `xorm:"not null comment('手机号') unique BIGINT"`
	Email           string    `xorm:"not null comment('账号名称') unique VARCHAR(255)"`
	EmailVerifiedAt time.Time `xorm:"TIMESTAMP"`
	Password        string    `xorm:"not null VARCHAR(255)"`
	RememberToken   string    `xorm:"VARCHAR(100)"`
	LastToken       string    `xorm:"TEXT"`
	Status          int       `xorm:"not null default 1 INT"`
	CreatedAt       time.Time `xorm:"TIMESTAMP"`
	UpdatedAt       time.Time `xorm:"TIMESTAMP"`
}

func (m *Users) TableName() string {
	return "users"
}
