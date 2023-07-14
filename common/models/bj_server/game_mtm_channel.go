package bj_server

type GameMtmChannel struct {
	Id        uint   `xorm:"not null pk autoincr UNSIGNED INT"`
	GameId    uint   `xorm:"not null default 0 comment('游戏ID') unique(uk_gid_cid) UNSIGNED SMALLINT"`
	ChannelId uint   `xorm:"not null default 0 comment('渠道ID') unique(uk_gid_cid) UNSIGNED INT"`
	Display   int    `xorm:"not null default 1 comment('1显示，0隐藏') SMALLINT"`
	Lang      string `xorm:"not null default 'en,ve' VARCHAR(36)"`
}

func (m *GameMtmChannel) TableName() string {
	return "game_mtm_channel"
}
