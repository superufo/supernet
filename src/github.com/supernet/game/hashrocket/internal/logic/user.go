package logic

type User struct {
	UserName     string
	BetAccount   float64 //下注大小
	EscapePoint  float64 // 逃生点
	reward       float64 //获得奖励
	BetRanker    int     //投注排名
	RewardRanker int     //奖励排名
}
