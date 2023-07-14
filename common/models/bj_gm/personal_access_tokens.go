package bj_gm

import (
	"time"
)

type PersonalAccessTokens struct {
	Id            uint64    `xorm:"not null pk autoincr UNSIGNED BIGINT"`
	TokenableType string    `xorm:"not null index(personal_access_tokens_tokenable_type_tokenable_id_index) VARCHAR(255)"`
	TokenableId   uint64    `xorm:"not null index(personal_access_tokens_tokenable_type_tokenable_id_index) UNSIGNED BIGINT"`
	Name          string    `xorm:"not null VARCHAR(255)"`
	Token         string    `xorm:"not null unique VARCHAR(64)"`
	Abilities     string    `xorm:"TEXT"`
	LastUsedAt    time.Time `xorm:"TIMESTAMP"`
	CreatedAt     time.Time `xorm:"TIMESTAMP"`
	UpdatedAt     time.Time `xorm:"TIMESTAMP"`
}

func (m *PersonalAccessTokens) TableName() string {
	return "personal_access_tokens"
}
