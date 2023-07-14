package bj_gm

import (
	"time"
)

type ActivityLog struct {
	Id          uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT"`
	LogName     string    `xorm:"index VARCHAR(255)"`
	Description string    `xorm:"not null TEXT"`
	SubjectType string    `xorm:"index(subject) VARCHAR(255)"`
	Event       string    `xorm:"VARCHAR(255)"`
	SubjectId   uint64    `xorm:"index(subject) UNSIGNED BIGINT"`
	CauserType  string    `xorm:"index(causer) VARCHAR(255)"`
	CauserId    uint64    `xorm:"index(causer) UNSIGNED BIGINT"`
	Properties  string    `xorm:"JSON"`
	Ip          int64     `xorm:"not null index BIGINT"`
	BatchUuid   string    `xorm:"CHAR(36)"`
	CreatedAt   time.Time `xorm:"not null TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"TIMESTAMP"`
}

func (m *ActivityLog) TableName() string {
	return "activity_log"
}
