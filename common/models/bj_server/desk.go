package bj_server

type Desk struct {
	Id       int    `xorm:"not null pk autoincr unique INT"`
	RoomId   int    `xorm:"not null comment('房间号码') unique INT"`
	GameId   int    `xorm:"not null comment('游戏id') INT"`
	Status   int    `xorm:"comment('状态 空闲 坐满 等') INT"`
	Players  string `xorm:"comment('玩家列表') VARCHAR(10000)"`
	Audience string `xorm:"comment('观众列表') VARCHAR(10000)"`
	Options  string `xorm:"comment('其他属性') JSON"`
}

func (m *Desk) TableName() string {
	return "desk"
}
