package bj_server

type Users struct {
	SId              string `xorm:"s_id not null pk VARCHAR(32)" json:"s_id" redis:"s_id"`
	Id               int    `xorm:"id INT" json:"id" redis:"id"`
	Account          string `xorm:"account VARCHAR(128)" json:"account" redis:"account"`
	Name             string `xorm:"name  VARCHAR(128)" json:"name" redis:"name"`
	Token            string `xorm:"token  VARCHAR(128)" json:"token" redis:"token"`
	Platform         string `xorm:"platform  comment('渠道标识') VARCHAR(32)" json:"platform" redis:"platform"`
	Sex              int    `xorm:"sex  TINYINT" json:"sex" redis:"sex"`
	Mac              string `xorm:"mac  comment('快速登录的账号') VARCHAR(64)" json:"mac" redis:"mac"`
	Nickname         string `xorm:"nickname  VARCHAR(128)" json:"nickname" redis:"nickname"`
	CCode            string `xorm:"c_code  comment('国家代码') VARCHAR(16)" json:"c_code" redis:"c_code"`
	Phone            string `xorm:"phone  VARCHAR(32)" json:"phone" redis:"phone"`
	RegisterTime     int    `xorm:"register_time  INT" json:"register_time" redis:"register_time"`
	Password         string `xorm:"password  comment('密码(游客米有MD5)') VARCHAR(128)" json:"password" redis:"password"`
	Agent            string `xorm:"agent  default '0' comment('代理标识') VARCHAR(128)" json:"agent" redis:"agent"`
	Status           int    `xorm:"status  default 0 comment('用户状态0正常1冻结') TINYINT" json:"status" redis:"status"`
	RegisterIp       string `xorm:"register_ip  comment('注册ip') VARCHAR(32)" json:"register_ip" redis:"register_ip"`
	GrandfatherId    string `xorm:"grandfather_id  default '' comment('上上级') VARCHAR(32)" json:"grandfather_id" redis:"grandfather_id"`
	TotalDirectNum   uint   `xorm:"total_direct_num  not null default 0 comment('直属下级成员') INT" json:"total_direct_num" redis:"total_direct_num"`
	TotalOtherNum    uint   `xorm:"total_other_num  not null default 0 comment('其他成员(下级的直属成员)') INT" json:"total_other_num" redis:"total_other_num"`
	ZaloId           string `xorm:"zalo_id  default '' comment('zalo_id') VARCHAR(128)" json:"zalo_id" redis:"zalo_id"`
	FacebookId       string `xorm:"facebook_id  default '' comment('facebook_id') VARCHAR(128)" json:"facebook_id" redis:"facebook_id"`
	GoogleId         string `xorm:"google_id  default '' comment('google_id') VARCHAR(128)" json:"google_id" redis:"google_id"`
	Organic          uint   `xorm:"organic  not null default 2 comment('自然属性(1 广告 2 自然)') TINYINT" json:"organic" redis:"organic"`
	TokenTime        int    `xorm:"token_time  default 0 comment('token有效时间') INT" json:"token_time" redis:"token_time"`
	DeviceId         string `xorm:"device_id  comment('设备ID登陆') VARCHAR(64)" json:"device_id" redis:"device_id"`
	InvitationCode   string `xorm:"invitation_code  not null default '' comment('邀请码') VARCHAR(5)" json:"invitation_code" redis:"invitation_code"`
	RegisterDeviceId string `xorm:"register_device_id  not null default '' comment('注册的设备ID') VARCHAR(64)" json:"register_device_id" redis:"register_device_id"`
	FatherId         string `xorm:"father_id  default '' comment('上级') VARCHAR(32)" json:"father_id" redis:"father_id"`
	Path             string `xorm:"path  default '' comment('关系链') VARCHAR(2048)" json:"path" redis:"path"`
	AgentType        uint   `xorm:"agent_type  default 1 comment('代理模式:1三级代理2金字塔') TINYINT" json:"agent_type" redis:"agent_type"`
}

func (m *Users) TableName() string {
	return "users"
}
