package bj_log

type Log2PerRoundHonghei struct {
	PerRoundSid string `xorm:"not null pk VARCHAR(32)"`
	RoomId      int    `xorm:"SMALLINT"`
	DataTime    int    `xorm:"default 0 comment('记录时间') INT"`
	UsersData   string `xorm:"comment('参与的用户:[{sid,输赢值},....]0平') MEDIUMTEXT"`
	Result      string `xorm:"default '' comment('开奖结果') VARCHAR(64)"`
	ResultState int    `xorm:"default 0 comment('平台输赢') TINYINT"`
	Amount      int64  `xorm:"default 0 comment('平台输赢值') BIGINT"`
}

func (m *Log2PerRoundHonghei) TableName() string {
	return "log_2_per_round_honghei"
}
