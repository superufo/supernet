package bj_server

type UserInfo struct {
	SId                   string  `xorm:"s_id not null pk VARCHAR(32)" json:"s_id" redis:"s_id"`
	LoginTime             int     `xorm:"login_time INT" json:"login_time" redis:"login_time"`
	OfflineTime           int     `xorm:"offline_time INT" json:"offline_time" redis:"offline_time"`
	Gold                  int64   `xorm:"gold BIGINT" json:"gold" redis:"gold"`
	Diamonds              int64   `xorm:"diamonds BIGINT" json:"diamonds" redis:"diamonds"`
	State                 int     `xorm:"state default 0 comment('0不在线1在线') TINYINT" json:"state" redis:"state"`
	LoginIp               string  `xorm:"login_ip comment('登录的自己IP') VARCHAR(64)" json:"login_ip" redis:"login_ip"`
	LoginSFlag            string  `xorm:"login_s_flag comment('当前登录的服务器') VARCHAR(64)" json:"login_s_flag" redis:"login_s_flag"`
	CtrlStatus            int     `xorm:"ctrl_status default 0 comment('0不被控制2Gm放1Gm杀') TINYINT" json:"ctrl_status" redis:"ctrl_status"`
	GameId                int     `xorm:"game_id default 0 comment('当前所在的游戏') SMALLINT" json:"game_id" redis:"game_id"`
	RoomId                int     `xorm:"room_id default 0 comment('当前所在的房间') SMALLINT" json:"room_id" redis:"room_id"`
	DeskId                int     `xorm:"desk_id default 0 comment('桌子id') INT" json:"desk_id" redis:"desk_id"`
	CtrlValue             int64   `xorm:"ctrl_value default 0 comment('已控制的金额') BIGINT" json:"ctrl_value" redis:"ctrl_value"`
	PStock                int64   `xorm:"p_stock default 0 comment('个人库存') BIGINT" json:"p_stock" redis:"p_stock"`
	RecentPlayTime        int     `xorm:"recent_play_time default 0 comment('最近游戏的时间用来后台排序(j金币变化的时间)') INT" json:"recent_play_time" redis:"recent_play_time"`
	TotalRecharge         int64   `xorm:"total_recharge default 0 comment('累计充值') BIGINT" json:"total_recharge" redis:"total_recharge"`
	TotalCash             int64   `xorm:"total_cash default 0 comment('累计提现') BIGINT" json:"total_cash" redis:"total_cash"`
	GmAward1              int64   `xorm:"gm_award_1 default 0 comment('客服奖励') BIGINT" json:"gm_award_1" redis:"gm_award_1"`
	GmAward2              int64   `xorm:"gm_award_2 default 0 comment('客服补偿') BIGINT" json:"gm_award_2" redis:"gm_award_2"`
	RecentPlayPerRoundSid string  `xorm:"recent_play_per_round_sid default '' comment('上一局牌局') VARCHAR(32)" json:"recent_play_per_round_sid" redis:"recent_play_per_round_sid"`
	CtrlData              int64   `xorm:"ctrl_data default 0 comment('需要控制金额') BIGINT" json:"ctrl_data" redis:"ctrl_data"`
	CtrlProbability       int     `xorm:"ctrl_probability default 0 comment('控制概率万分比') INT" json:"ctrl_probability" redis:"ctrl_probability"`
	CtrlScales            float32 `xorm:"ctrl_scales default 0 comment('控制金额的浮动比例0.1') FLOAT" json:"ctrl_scales" redis:"ctrl_scales"`
	Platform              string  `xorm:"platform comment('渠道标识') VARCHAR(32)" json:"platform" redis:"platform"`
	Agent                 string  `xorm:"agent comment('代理标识') VARCHAR(32)" json:"agent" redis:"agent"`
}

func (m *UserInfo) TableName() string {
	return "user_info"
}
