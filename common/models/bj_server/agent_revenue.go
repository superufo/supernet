package bj_server

type AgentRevenue struct {
	Id          int    `xorm:"not null pk INT"`
	RevenueTime int    `xorm:"default 0 comment('收益时间') INT"`
	Change      int64  `xorm:"default 0 comment('收益数值') BIGINT"`
	RevenueFrom string `xorm:"comment('收益来源谁') VARCHAR(32)"`
	RevenueTo   string `xorm:"comment('收益给谁') VARCHAR(32)"`
	State       int    `xorm:"default 0 comment('0未分配1已分配') TINYINT"`
	StateTime   int    `xorm:"default 0 comment('状态改变的时间') INT"`
}

func (m *AgentRevenue) TableName() string {
	return "agent_revenue"
}
