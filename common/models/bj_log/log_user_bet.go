package bj_log

type LogUserBet struct {
	Id        uint   `xorm:"id not null pk autoincr  INT" json:"id" redis:"id"`
	ChannelId int    `xorm:"channel_id not null default 0 comment('渠道号') INT" json:"channel_id" redis:"channel_id"`
	SId       string `xorm:"s_id not null default '' comment('用户id') index VARCHAR(32)" json:"s_id" redis:"s_id"`
	GameId    int    `xorm:"game_id not null default 0 comment('游戏类型') SMALLINT" json:"game_id" redis:"game_id"`
	Bet       uint64 `xorm:"bet not null default 0 comment('有效下注金额')   BIGINT" json:"bet" redis:"bet"`
	CreatedAt uint   `xorm:"created_at not null comment('时间')  INT" json:"created_at" redis:"created_at"`
}

func (m *LogUserBet) TableName() string {
	return "log_user_bet"
}
