package bj_log

type LogUserPerRound struct {
	Id                  int    `xorm:"not null pk autoincr comment('id') INT"`
	UserSid             string `xorm:"VARCHAR(32)"`
	PerRoundSid         string `xorm:"VARCHAR(32)"`
	GameId              int    `xorm:"default 1 comment('玩法') TINYINT"`
	RoomId              int    `xorm:"default 0 comment('房间') TINYINT"`
	Change              int64  `xorm:"default 0 comment('这局收入') BIGINT"`
	EndTime             int    `xorm:"default 0 comment('时间') INT"`
	Bets                string `xorm:"comment('下注情况json') VARCHAR(255)"`
	Result              string `xorm:"comment('开牌结果json') VARCHAR(255)"`
	PerRoundState       int    `xorm:"default 0 comment('当前牌局状态') TINYINT"`
	Win                 int64  `xorm:"default 0 comment('输赢值') BIGINT"`
	BeforeMoney         int64  `xorm:"comment('原来多少') BIGINT"`
	AfterMoney          int64  `xorm:"comment('结算之后多少') BIGINT"`
	Platform            string `xorm:"comment('渠道') VARCHAR(32)"`
	Agent               string `xorm:"comment('代理') VARCHAR(32)"`
	PlayerServiceCharge int64  `xorm:"not null default 0 comment('玩家服务费') BIGINT"`
}

func (m *LogUserPerRound) TableName() string {
	return "log_user_per_round"
}
