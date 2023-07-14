package bj_log

type LogLoginLog struct {
	Id        int    `xorm:"not null pk INT"`
	SId       string `xorm:"comment('人物唯一标识') VARCHAR(32)"`
	LoginTime int    `xorm:"INT"`
	LoginIp   string `xorm:"comment('登录ip') VARCHAR(64)"`
}

func (m *LogLoginLog) TableName() string {
	return "log_login_log"
}
