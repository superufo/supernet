package bj_server

type AgentInfo struct {
	SId     string `xorm:"not null pk VARCHAR(32)"`
	Fathers string `xorm:"VARCHAR(255)"`
	Config  string `xorm:"VARCHAR(255)"`
}

func (m *AgentInfo) TableName() string {
	return "agent_info"
}
