package bj_server

type CarouselConfig struct {
	Id        uint   `xorm:"not null pk autoincr UNSIGNED INT"`
	ChannelId int    `xorm:"not null comment('渠道id') index INT"`
	Type      int    `xorm:"not null default 1 comment('类型:1纯图片2跳转地址3跳转游戏') TINYINT"`
	Poster    string `xorm:"not null comment('海报地址') VARCHAR(200)"`
	Link      string `xorm:"comment('链接') VARCHAR(200)"`
	CreatedAt uint   `xorm:"not null comment('创建时间') UNSIGNED INT"`
	UpdatedAt uint   `xorm:"not null comment('更新时间') UNSIGNED INT"`
}

func (m *CarouselConfig) TableName() string {
	return "carousel_config"
}
