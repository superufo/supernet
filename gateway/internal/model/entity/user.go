package entity

type Users struct {
	SId          string `xorm:"s_id" json:"s_id" redis:"s_id"`
	ID           int    `xorm:"id" json:"id" redis:"id"`
	Account      string `xorm:"account" json:"account" redis:"account"`
	Name         string `xorm:"name" json:"name" redis:"name"`
	Token        string `xorm:"token" json:"token" redis:"token"`
	Platform     string `xorm:"platform" json:"platform" redis:"platform"`
	Sex          int8   `xorm:"sex" json:"sex" redis:"sex"`
	Mac          string `xorm:"mac" json:"mac" redis:"mac"`
	Nickname     string `xorm:"nickname" json:"nickname" redis:"nickname"`
	CCode        string `xorm:"c_code" json:"c_code" redis:"c_code"`
	Phone        string `xorm:"phone" json:"phone" redis:"phone"`
	RegisterTime int64  `xorm:"register_time" json:"register_time" redis:"register_time"`
	Password     string `xorm:"password" json:"password" redis:"password"`
	Agent        string `xorm:"agent" json:"agent" redis:"agent"`
	Status       int8   `xorm:"status" json:"status" redis:"status"`
	RegisterIp   string `xorm:"register_ip" json:"register_ip" redis:"register_ip"`
	FatherId     string `xorm:"father_id" json:"father_id" redis:"father_id"`
}

const TABLE_USERS = "users"
