package bj_gm

import (
	"time"
)

type Users struct {
	Id              uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT"`
	Name            string    `xorm:"not null VARCHAR(255)"`
	Avatar          string    `xorm:"not null VARCHAR(255)"`
	Email           string    `xorm:"not null unique VARCHAR(255)"`
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
