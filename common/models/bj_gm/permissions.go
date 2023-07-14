package bj_gm

import (
	"time"
)

type Permissions struct {
	Id         uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT"`
	ParentId   int64     `xorm:"not null default 0 BIGINT"`
	Name       string    `xorm:"not null unique(permissions_name_guard_name_unique) VARCHAR(255)"`
	Title      string    `xorm:"not null VARCHAR(255)"`
	Path       string    `xorm:"not null VARCHAR(255)"`
	Component  string    `xorm:"VARCHAR(255)"`
	Sort       int       `xorm:"not null default 1 INT"`
	Hidden     int       `xorm:"not null default 0 INT"`
	Type       string    `xorm:"not null default 'parent' ENUM('button','children','parent')"`
	Icon       string    `xorm:"VARCHAR(255)"`
	AlwaysShow int       `xorm:"not null default 0 INT"`
	Redirect   string    `xorm:"VARCHAR(255)"`
	GuardName  string    `xorm:"not null unique(permissions_name_guard_name_unique) VARCHAR(255)"`
	CreatedAt  time.Time `xorm:"TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"TIMESTAMP"`
}

func (m *Permissions) TableName() string {
	return "permissions"
}
