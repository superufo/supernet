package def

//流水类型定义[1,2,3,....]
const (
	CHANGE_BET         uint8 = 1 + iota //1-下注
	CHANGE_WINLOSE                      //2-游戏输赢
	CHANGE_ADMINADD                     //3-后台加钱
	CHANGE_ADMINREDUCE                  //4-后台减钱
	CHANGE_AGENT                        //5-代理
	CHANGE_REWARD                       //6-奖金
	CHANGE_ACTIVE                       //7-活动
)
