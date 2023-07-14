package bj_server

type GamesConfig struct {
	Id                  int     `xorm:"id not null pk INT" json:"id" redis:"id"`
	Name                string  `xorm:"name VARCHAR(255)" json:"name" redis:"name"`
	Stock1              int64   `xorm:"stock_1 default 0 comment('库存1基础库存') BIGINT" json:"stock_1" redis:"stock_1"`
	Stock1WarnWater     int64   `xorm:"stock_1_warn_water default 0 comment('库存1警戒值') BIGINT" json:"stock_1_warn_water" redis:"stock_1_warn_water"`
	DrawWater           int     `xorm:"draw_water default 0 comment('库存1目标和报警之间的抽水概几率万分比') INT" json:"draw_water" redis:"draw_water"`
	PlayerServiceCharge float32 `xorm:"player_service_charge default 0 comment('玩家赢的玩家服务费') FLOAT" json:"player_service_charge" redis:"player_service_charge"`
	SystemServiceCharge float32 `xorm:"system_service_charge default 0 comment('玩家赢的系统服务费(暂时没用到)') FLOAT" json:"system_service_charge" redis:"system_service_charge"`
	Stock2ServiceCharge float32 `xorm:"stock_2_service_charge comment('库存2奖励库存比例') FLOAT" json:"stock_2_service_charge" redis:"stock_2_service_charge"`
	Stock2WarnWater     int64   `xorm:"stock_2_warn_water comment('库存2报警值') BIGINT" json:"stock_2_warn_water" redis:"stock_2_warn_water"`
	Stock1State         int     `xorm:"stock_1_state default 0 comment('库存1状态') TINYINT" json:"stock_1_state" redis:"stock_1_state"`
	UpdateTime          int     `xorm:"update_time default 0 comment('变化时间') INT" json:"update_time" redis:"update_time"`
	ToStock1            float32 `xorm:"to_stock_1 comment('玩家输了,钱进库存的比例') FLOAT" json:"to_stock_1" redis:"to_stock_1"`
	GameId              int     `xorm:"game_id comment('关联的游戏id') INT" json:"game_id" redis:"game_id"`
}

func (m *GamesConfig) TableName() string {
	return "games_config"
}
