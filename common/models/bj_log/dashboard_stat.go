package bj_log

import (
	"time"
)

type DashboardStat struct {
	Id               int       `xorm:"not null pk autoincr INT"`
	StatisticTime    time.Time `xorm:"not null comment('统计时间') DATE"`
	RegisterCnt      int64     `xorm:"not null default 0 comment('注册人数') BIGINT"`
	PlayerBetCnt     int64     `xorm:"not null default 0 comment('玩家投注量') BIGINT"`
	AgentCnt         int64     `xorm:"not null default 0 comment('代理人数') BIGINT"`
	WithdrawCnt      int64     `xorm:"not null default 0 comment('提现人数') BIGINT"`
	WithdrawMoneyCnt int64     `xorm:"not null default 0 comment('提现金额') BIGINT"`
	RechargeCnt      int64     `xorm:"not null default 0 comment('充值人数') BIGINT"`
	RechargeMoneyCnt int64     `xorm:"not null default 0 comment('充值金额') BIGINT"`
	ActiveCnt        int64     `xorm:"not null default 0 comment('活跃人数') BIGINT"`
	ServiceMoneyCnt  int64     `xorm:"not null default 0 comment('服务费统计') BIGINT"`
	FlowWaterCnt     int64     `xorm:"not null default 0 comment('流水统计') BIGINT"`
	Game1Cnt         int64     `xorm:"not null default 0 comment('game_1:龙虎斗') BIGINT"`
	Game2Cnt         int64     `xorm:"not null default 0 comment('game_2:红黑') BIGINT"`
	Game3Cnt         int64     `xorm:"not null default 0 comment('game_3:百家乐') BIGINT"`
	Game4Cnt         int64     `xorm:"not null default 0 comment('game_4:鱼虾蟹') BIGINT"`
	ChannelId        int       `xorm:"not null default 0 comment('渠道号') INT"`
	CreatedAt        time.Time `xorm:"not null DATETIME"`
}

func (m *DashboardStat) TableName() string {
	return "dashboard_stat"
}
