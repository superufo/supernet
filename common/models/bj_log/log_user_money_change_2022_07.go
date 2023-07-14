package bj_log

type LogUserMoneyChange202207 struct {
	Id          int    `xorm:"not null pk autoincr INT"`
	SId         string `xorm:"comment('用户唯一id') VARCHAR(32)"`
	MoneyType   int    `xorm:"default 0 comment('1金币2钻石') TINYINT"`
	BeforeMoney int64  `xorm:"comment('原来多少钱') BIGINT"`
	ChangeType  int    `xorm:"default 0 comment('变化原因类型1,2,3,4,5,6,7,8,9') TINYINT"`
	Change      int64  `xorm:"default 0 comment('变化了多少有正有负') BIGINT"`
	AfterMoney  int64  `xorm:"default 0 comment('变化之后剩余多少') BIGINT"`
	PerRoundSid string `xorm:"comment('牌局号') VARCHAR(64)"`
	GameId      int    `xorm:"default 0 comment('游戏id 0非游戏的') TINYINT"`
	RoomId      int    `xorm:"default 0 comment('房间id') TINYINT"`
	UpTime      int    `xorm:"default 0 comment('改变的时间') INT"`
	SerialNo    string `xorm:"VARCHAR(32)"`
	Platform    string `xorm:"comment('渠道') VARCHAR(32)"`
	Agent       string `xorm:"comment('代理') VARCHAR(32)"`
}

func (m *LogUserMoneyChange202207) TableName() string {
	return "log_user_money_change_2022_07"
}
