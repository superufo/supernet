package msg

// 所有消息类型  消息为两位
const (
	CMD_USER        uint16 = 800 //用户信息
	CMD_USER_INFO   uint16 = 801 // 玩家金币等信息
	CMD_GAME_CONFIG uint16 = 802 // 游戏配置
	CMD_LOGIN       uint16 = 803 // 登录

	CMD_LOG uint16 = 804 // 日志信息

	CMDS string = "800,801,802,803,804"
)
