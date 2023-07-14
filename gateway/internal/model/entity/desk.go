package entity

type Desk struct {
	ID       int64  `gorm:"column:id" xorm:"id" json:"id"`
	RoomId   int64  `gorm:"column:room_id" xorm:"room_id" json:"room_id"`
	GameId   int64  `gorm:"column:game_id" xorm:"game_id" json:"game_id"`
	Status   int64  `gorm:"column:status" xorm:"status" json:"status"`
	Players  string `gorm:"column:players" xorm:"players" json:"players"`
	Audience string `gorm:"column:audience" xorm:"audience" json:"audience"`
}

const TABLE_DESK = "desk"
