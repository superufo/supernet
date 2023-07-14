package enum

// 所有消息类型  消息为两位
const (
	CMD_READY             uint16 = 5100
	CMD_ENTER_GAME        uint16 = 5110 //          客户端进入游戏
	CMD_DOING_INFO        uint16 = 5111 //    进入游戏,回复当前信息
	CMD_BET               uint16 = 5112 //          上傳的投注信息
	CMD_BET_10            uint16 = 5113 //      排名前10的投注信息
	CMD_EXPLOSION         uint16 = 5114 //        发送爆点给客户端
	CMD_PAST100_EXPLOSION uint16 = 5115 //发送过去100局的爆点给客户
)

// 消息的合集
const (
	CMDS string = "5100,5101,5103,5104,5105"
)

// 定義狀態
type (
	State int
)

// 游戏状态 state
const (
	READY int = iota + 1
	BETING
	RUNNING
	END
)
