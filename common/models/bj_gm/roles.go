package bj_gm

import (
	"time"
)

type Roles struct {
	Id          uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT"`
	Name        string    `xorm:"not null unique(roles_name_guard_name_unique) VARCHAR(255)"`
	GuardName   string    `xorm:"not null unique(roles_name_guard_name_unique) VARCHAR(255)"`
	Description string    `xorm:"not null VARCHAR(255)"`
	Routes      string    `xorm:"not null JSON"`
	CreatedAt   time.Time `xorm:"TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"TIMESTAMP"`
}

func (m *Roles) TableName() string {
	return "roles"
}
