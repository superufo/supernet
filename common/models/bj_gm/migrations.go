package bj_gm

type Migrations struct {
	Id        uint   `xorm:"not null pk autoincr UNSIGNED INT"`
	Migration string `xorm:"not null VARCHAR(255)"`
	Batch     int    `xorm:"not null INT"`
}

func (m *Migrations) TableName() string {
	return "migrations"
}
