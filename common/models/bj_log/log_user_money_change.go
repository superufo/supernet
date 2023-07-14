package bj_log

type LogUserMoneyChange struct {
	Id          int    `xorm:"id not null pk autoincr INT" json:"id" redis:"id"`
	SId         string `xorm:"s_id comment('用户唯一id') VARCHAR(32)" json:"s_id" redis:"s_id"`
	MoneyType   int    `xorm:"money_type default 0 comment('1金币2钻石') TINYINT" json:"money_type" redis:"money_type"`
	BeforeMoney int64  `xorm:"before_money comment('原来多少钱') BIGINT" json:"before_money" redis:"before_money"`
	ChangeType  int    `xorm:"change_type default 0 comment('变化原因类型1,2,3,4,5,6,7,8,9') TINYINT" json:"change_type" redis:"change_type"`
	Change      int64  `xorm:"change default 0 comment('变化了多少有正有负') BIGINT" json:"change" redis:"change"`
	AfterMoney  int64  `xorm:"after_money default 0 comment('变化之后剩余多少') BIGINT" json:"after_money" redis:"after_money"`
	PerRoundSid string `xorm:"per_round_sid comment('牌局号') VARCHAR(64)" json:"per_round_sid" redis:"per_round_sid"`
	GameId      int    `xorm:"game_id default 0 comment('游戏id 0非游戏的') TINYINT" json:"game_id" redis:"game_id"`
	RoomId      int    `xorm:"room_id default 0 comment('房间id') TINYINT" json:"room_id" redis:"room_id"`
	UpTime      int    `xorm:"up_time default 0 comment('改变的时间') INT" json:"up_time" redis:"up_time"`
	SerialNo    string `xorm:"serial_no comment('流水号') VARCHAR(32)" json:"serial_no" redis:"serial_no"`
	Platform    string `xorm:"platform comment('渠道') VARCHAR(32)" json:"platform" redis:"platform"`
	Agent       string `xorm:"agent comment('代理') VARCHAR(32)" json:"agent" redis:"agent"`
}

func (m *LogUserMoneyChange) TableName() string {
	return "log_user_money_change"
}
