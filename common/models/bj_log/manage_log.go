package bj_log

type ManageLog struct {
	Id   int    `xorm:"not null pk INT"`
	Name string `xorm:"comment('日志名称') VARCHAR(128)"`
}

func (m *ManageLog) TableName() string {
	return "manage_log"
}
