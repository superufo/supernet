package entity

import "time"

// GameInstance undefined
type GameInstance struct {
	GameSid   string    `json:"game_sid" xorm:"game_sid"`
	Starttime time.Time `json:"starttime"`
	Explosion float64   `json:"explosion"`
	EndTime   time.Time `json:"end_time"`
}

// TableName 表名称
func (*GameInstance) TableName() string {
	return "game_instance"
}

//const TABLE_USERS = "game_instance"
