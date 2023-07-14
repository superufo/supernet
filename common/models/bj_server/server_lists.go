package bj_server

type ServerLists struct {
	ServerFlag        string `xorm:"not null pk comment('服务器标识唯一') VARCHAR(255)"`
	ServerAddr        string `xorm:"comment('服务器地址') VARCHAR(255)"`
	ServerWebsockPort int    `xorm:"default 0 comment('游戏端口') INT"`
	ServerWebPort     int    `xorm:"default 0 comment('web端口') INT"`
	State             int    `xorm:"default 0 comment('0关1开') TINYINT"`
	Node              string `xorm:"comment('节点信息') VARCHAR(255)"`
	OnlineNum         int    `xorm:"default 0 comment('在线人数') INT"`
}

func (m *ServerLists) TableName() string {
	return "server_lists"
}
