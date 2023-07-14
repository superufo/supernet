package msg

// 所有消息类型  消息为两位
const (
	CMD_SERVER_LIST           uint16 = 3016 // 下发服务信息
	CMD_STEP                  uint16 = 3003 // 下发步长
	CMD_LEAVE_GAME            uint16 = 2004
	CMD_GAME_1_DESK           uint16 = 3002
	CMD_GAME_1_BETS           uint16 = 3005
	CMD_GAME_1_CHANGE_STATE   uint16 = 3006
	CMD_GAME_1_END_RESULT     uint16 = 3007
	CMD_GAME_1_START_NEW_BETS uint16 = 3008
	CMD_GAME_1_END_AWARD      uint16 = 3009
	CMD_NOTICE_MSG            uint16 = 1004

	CMD_ERROR     uint16 = 99
	CMD_HEART_BIT uint16 = 777 // 心跳信息

	CMD_USER        uint16 = 800 //用户信息
	CMD_USER_INFO   uint16 = 801 // 玩家金币等信息
	CMD_GAME_CONFIG uint16 = 802 // 游戏配置
	CMD_LOGIN       uint16 = 803 // 登录
	CMD_REG         uint16 = 805 // 注册

	CMD_LOG uint16 = 804 // 日志信息

	CMD_CLOSE uint16 = 1111 // 断线信息

	CMD_GAME_2_SESSION       uint16 = 2102 // 选择场次,回复该场次内信息
	CMD_Game_2_Eeter_Desk    uint16 = 2103 // 玩家进入桌子
	CMD_GAME_2_BETS          uint16 = 2104 // 玩家投注,返回投注的信息
	CMD_GAME_2_END_RESULT    uint16 = 2105 // 游戏结束的结果
	CMD_GAME_2_OUT           uint16 = 2106 // 推出游戏
	CMD_GAME_2_STATUS_CHANGE uint16 = 2107 // 状态改变
	CMD_GAME_GET_RECORD      uint16 = 2108 // 获取游戏记录

	CMDS string = "800,801,802,803,804,1001,1004,1005,1006,1010,2001,2003,2004,3002,3004,3005,3006,3007,3008,3009,3001,3003,3004,3007,3016"
)
