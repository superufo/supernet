package bj_server

type GamesInfo struct {
	Id            int   `xorm:"id not null pk TINYINT" json:"id" redis:"id"`
	CurrentStock1 int64 `xorm:"current_stock_1 default 0 comment('当前水位') BIGINT" json:"current_stock_1" redis:"current_stock_1"`
	CurrentStock2 int64 `xorm:"current_stock_2 comment('当前库存2水位') BIGINT" json:"current_stock_2" redis:"current_stock_2"`
	ChangeTime    int   `xorm:"change_time default 0 comment('当前水位变化的时间') INT" json:"change_time" redis:"change_time"`
	UpdateTime    int   `xorm:"update_time default 0 comment('写入数据库的时间') INT" json:"update_time" redis:"update_time"`
	GameId        int   `xorm:"game_id comment('关联的游戏id') SMALLINT" json:"game_id" redis:"game_id"`
}

func (m *GamesInfo) TableName() string {
	return "games_info"
}
