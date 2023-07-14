package bj_server

import (
	"time"
)

type Room struct {
	Id            uint      `xorm:"id not null pk autoincr unique INT" json:"id" redis:"id"` 
	Name          string    `xorm:"name comment('房间名称如高级体验场') VARCHAR(45)" json:"name" redis:"name"`
	Score         int64     `xorm:"score comment('准入条件 金币数') BIGINT" json:"score" redis:"score"`
	GameId        int       `xorm:"game_id comment('游戏id') INT" json:"game_id" redis:"game_id"`
	MaxBet        int64     `xorm:"max_bet comment('最大下注') BIGINT" json:"max_bet" redis:"max_bet"`
	MinBet        int64     `xorm:"min_bet comment('最小下注') BIGINT" json:"min_bet" redis:"min_bet"`
	MaxBetLucky   int64     `xorm:"max_bet_lucky comment('幸运一击的最大值') BIGINT" json:"max_bet_lucky" redis:"max_bet_lucky"`
	MinBetLucky   int64     `xorm:"min_bet_lucky comment('幸运一击的最小值') BIGINT" json:"min_bet_lucky" redis:"min_bet_lucky"`
	Status        int       `xorm:"status comment('状态') INT" json:"status" redis:"status"`
	Remark        string    `xorm:"remark VARCHAR(100)" json:"remark" redis:"remark"`
	NewColumn     int       `xorm:"new_column INT" json:"new_column" redis:"new_column"`
	EnableRobot   uint      `xorm:"enable_robot not null default 0 comment('是否开放机器人') TINYINT" json:"enable_robot" redis:"enable_robot"`
	EnableIpLimit uint      `xorm:"enable_ip_limit not null default 0 comment('是否开启IP限制匹配') TINYINT" json:"enable_ip_limit" redis:"enable_ip_limit"`
	BetAreaLimit  string    `xorm:"bet_area_limit not null default '' comment('下注区域限红配置') VARCHAR(500)" json:"bet_area_limit" redis:"bet_area_limit"`
	MinScore      int64     `xorm:"min_score not null comment('准入条件 最小准入金币数') BIGINT" json:"min_score" redis:"min_score"`
	MaxScore      int64     `xorm:"max_score not null comment('准入条件 最大准入金币数') BIGINT" json:"max_score" redis:"max_score"`
	CreatedAt     time.Time `xorm:"created_at not null default CURRENT_TIMESTAMP TIMESTAMP" json:"created_at" redis:"created_at"`
	UpdatedAt     time.Time `xorm:"updated_at not null default CURRENT_TIMESTAMP TIMESTAMP" json:"updated_at" redis:"updated_at"`
}

func (m *Room) TableName() string {
	return "room"
}
