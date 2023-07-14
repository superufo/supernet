package entity

type ServerLists struct {
	ServerFlag        string `xorm:"server_flag" json:"server_flag"`
	ServerAddr        string `xorm:"server_addr" json:"server_addr"`
	ServerWebsockPort int64  `xorm:"server_websock_port" json:"server_websock_port"`
	ServerWebPort     int64  `xorm:"server_web_port" json:"server_web_port"`
	State             int64  `xorm:"state" json:"state"`
	Node              string `xorm:"node" json:"node"`
	OnlineNum         int64  `xorm:"online_num" json:"online_num"`
}

const TABLE_SERVER_LIST = "server_lists"
