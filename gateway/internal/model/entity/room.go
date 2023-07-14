package entity

type Room struct {
	ID          int64  `gorm:"column:id" xorm:"id" json:"id"`
	Name        string `gorm:"column:name" xorm:"name" json:"name"`
	Score       int64  `gorm:"column:score" xorm:"score" json:"score"`
	GameId      int64  `gorm:"column:game_id" xorm:"game_id" json:"game_id"`
	MaxBet      int64  `gorm:"column:max_bet" xorm:"max_bet" json:"max_bet"`
	MinBet      int64  `gorm:"column:min_bet" xorm:"min_bet" json:"min_bet"`
	MaxBetLucky int64  `gorm:"column:max_bet_lucky" xorm:"max_bet_lucky" json:"max_bet_lucky"`
	MinBetLucky int64  `gorm:"column:min_bet_lucky" xorm:"min_bet_lucky" json:"min_bet_lucky"`
	Remark      string `gorm:"column:remark" xorm:"remark" json:"remark"`
}

const TABLE_ROOM = "room"
