package bj_server

type Channels struct {
	ChannelId         uint    `xorm:"channel_id not null pk comment('渠道id') UNSIGNED INT" json:"channel_id" redis:"channel_id"`
	CLossCompanyRatio float32 `xorm:"c_loss_company_ratio not null comment('客损平台收益比例') FLOAT" json:"c_loss_company_ratio" redis:"c_loss_company_ratio"`
	CLossChannelRatio float32 `xorm:"c_loss_channel_ratio not null comment('客损渠道收益比例') FLOAT" json:"c_loss_channel_ratio" redis:"c_loss_channel_ratio"`
	DividendRatio     float32 `xorm:"dividend_ratio not null comment('渠道商与公司分成比例') FLOAT" json:"dividend_ratio" redis:"dividend_ratio"`
	LevelRatioLimit   float32 `xorm:"level_ratio_limit not null comment('代理分成上限') FLOAT" json:"level_ratio_limit" redis:"level_ratio_limit"`
	PyramidRatio      string  `xorm:"pyramid_ratio JSON" json:"pyramid_ratio" redis:"pyramid_ratio"`
	LevelRatio        string  `xorm:"level_ratio not null comment('代理分成比例') JSON" json:"level_ratio" redis:"level_ratio"`
	AgentType         int     `xorm:"agent_type not null default 1 comment('代理类型') TINYINT" json:"agent_type" redis:"agent_type"`
	AppKey            string  `xorm:"app_key not null comment('接入方app_key') VARCHAR(255)" json:"app_key" redis:"app_key"`
	AppId             string  `xorm:"app_id not null comment('接入方app_id') VARCHAR(255)" json:"app_id" redis:"app_id"`
}

func (m *Channels) TableName() string {
	return "channels"
}
