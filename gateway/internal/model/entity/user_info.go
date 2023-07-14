package entity

type UserInfo struct {
	SId                   string  `gorm:"column:s_id" xorm:"s_id" json:"s_id"`
	LoginTime             int64   `gorm:"column:login_time" xorm:"login_time" json:"login_time"`
	OfflineTime           int64   `gorm:"column:offline_time" xorm:"offline_time" json:"offline_time"`
	Gold                  int64   `gorm:"column:gold" xorm:"gold" json:"gold"`
	Diamonds              int64   `gorm:"column:diamonds" xorm:"diamonds" json:"diamonds"`
	State                 int64   `gorm:"column:state" xorm:"state" json:"state"`
	LoginIp               string  `gorm:"column:login_ip" xorm:"login_ip" json:"login_ip"`
	LoginSFlag            string  `gorm:"column:login_s_flag" xorm:"login_s_flag" json:"login_s_flag"`
	CtrlStatus            int64   `gorm:"column:ctrl_status" xorm:"ctrl_status" json:"ctrl_status"`
	GameId                int64   `gorm:"column:game_id" xorm:"game_id" json:"game_id"`
	RoomId                int64   `gorm:"column:room_id" xorm:"room_id" json:"room_id"`
	DeskId                int64   `gorm:"column:desk_id" xorm:"desk_id" json:"desk_id"`
	CtrlValue             int64   `gorm:"column:ctrl_value" xorm:"ctrl_value" json:"ctrl_value"`
	PStock                int64   `gorm:"column:p_stock" xorm:"p_stock" json:"p_stock"`
	RecentPlayTime        int64   `gorm:"column:recent_play_time" xorm:"recent_play_time" json:"recent_play_time"`
	TotalRecharge         int64   `gorm:"column:total_recharge" xorm:"total_recharge" json:"total_recharge"`
	TotalCash             int64   `gorm:"column:total_cash" xorm:"total_cash" json:"total_cash"`
	GmAward1              int64   `gorm:"column:gm_award_1" xorm:"gm_award_1" json:"gm_award_1"`
	GmAward2              int64   `gorm:"column:gm_award_2" xorm:"gm_award_2" json:"gm_award_2"`
	RecentPlayPerRoundSid string  `gorm:"column:recent_play_per_round_sid" xorm:"recent_play_per_round_sid" json:"recent_play_per_round_sid"`
	CtrlData              int64   `gorm:"column:ctrl_data" xorm:"ctrl_data" json:"ctrl_data"`
	CtrlProbability       int64   `gorm:"column:ctrl_probability" xorm:"ctrl_probability" json:"ctrl_probability"`
	CtrlScales            float64 `gorm:"column:ctrl_scales" xorm:"ctrl_scales" json:"ctrl_scales"`
	Platform              string  `gorm:"column:platform" xorm:"platform" json:"platform"`
	Agent                 string  `gorm:"column:agent" xorm:"agent" json:"agent"`
}

const TABLE_USER_INFO = "user_info"
