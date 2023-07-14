package entity

type Users struct {
	SId          string `gorm:"column:s_id" xorm:"s_id" json:"s_id"`
	ID           int64  `gorm:"column:id" xorm:"id" json:"id"`
	Account      string `gorm:"column:account" xorm:"account" json:"account"`
	Name         string `gorm:"column:name" xorm:"name" json:"name"`
	Token        string `gorm:"column:token" xorm:"token" json:"token"`
	Platform     string `gorm:"column:platform" xorm:"platform" json:"platform"`
	Sex          int64  `gorm:"column:sex" xorm:"sex" json:"sex"`
	Mac          string `gorm:"column:mac" xorm:"mac" json:"mac"`
	Nickname     string `gorm:"column:nickname" xorm:"nickname" json:"nickname"`
	CCode        string `gorm:"column:c_code" xorm:"c_code" json:"c_code"`
	Phone        string `gorm:"column:phone" xorm:"phone" json:"phone"`
	RegisterTime int64  `gorm:"column:register_time" xorm:"register_time" json:"register_time"`
	Password     string `gorm:"column:password" xorm:"password" json:"password"`
	Agent        string `gorm:"column:agent" xorm:"agent" json:"agent"`
	Status       int64  `gorm:"column:status" xorm:"status" json:"status"`
	RegisterIp   string `gorm:"column:register_ip" xorm:"register_ip" json:"register_ip"`
	FatherId     string `gorm:"column:father_id" xorm:"father_id" json:"father_id"`
}

const TABLE_USERS = "users"
