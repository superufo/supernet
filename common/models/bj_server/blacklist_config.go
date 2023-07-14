package bj_server

type BlacklistConfig struct {
	Id        uint   `xorm:"not null pk autoincr UNSIGNED INT"`
	Type      int    `xorm:"not null default 1 comment('类型:1 ip 2机器码') TINYINT"`
	Value     string `xorm:"not null comment('ip或机器码') index VARCHAR(255)"`
	Status    int    `xorm:"not null default 1 comment('状态:1正常2封禁') TINYINT"`
	Remarks   string `xorm:"comment('备注') VARCHAR(255)"`
	AdminId   int    `xorm:"not null comment('管理员id') INT"`
	CreatedAt uint   `xorm:"not null comment('创建时间') UNSIGNED INT"`
	UpdatedAt uint   `xorm:"not null comment('更新时间') UNSIGNED INT"`
}

func (m *BlacklistConfig) TableName() string {
	return "blacklist_config"
}
