package bj_log

import (
	"time"
)

type AgentProfitStat struct {
	Id         uint      `xorm:"id not null pk autoincr INT" json:"id" redis:"id"`
	ChannelId  int       `xorm:"channel_id not null comment('渠道号') INT" json:"channel_id" redis:"channel_id"`
	GameId     int       `xorm:"game_id not null comment('收益来源游戏') INT" json:"game_id" redis:"game_id"`
	ProfitFrom string    `xorm:"profit_from not null comment('收益来源ID') VARCHAR(32)" json:"profit_from" redis:"profit_from"`
	ProfitTo   string    `xorm:"profit_to not null comment('收益ID') VARCHAR(32)" json:"profit_to" redis:"profit_to"`
	Profit     int64     `xorm:"profit not null comment('收益金额') BIGINT" json:"profit" redis:"profit"`
	ProfitTime time.Time `xorm:"profit_time not null comment('收益计算时间') DATETIME" json:"profit_time" redis:"profit_time"`
}

func (m *AgentProfitStat) TableName() string {
	return "agent_profit_stat"
}
