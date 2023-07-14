package state

import (
	"github.com/supernet/game/hashrocket/enum"
	"github.com/supernet/game/hashrocket/internal/server"
	"github.com/supernet/game/hashrocket/pkg/log"

	"go.uber.org/zap"
	"time"
)

// GameInst 游戏的每一局定义
type GameInst struct {
	GameSid   string    // 本局的唯一标志
	startTime time.Time //游戏开始的时间戳
	EndTime   time.Time //游戏开始的时间戳
	Explosion float64   // 爆点
	EnterNum  int64     // 参与人数
	BetNum    int64     // 下注人数

	c  ifstate                // 当前状态
	ss map[enum.State]ifstate // 所有经历过的状态数据保存在内存

	isEnd bool //游戏结束
}

func NewGameInst() (*GameInst, error) {
	//生成 game_sid
	gis := server.NewGameInstanceService()
	g, err := gis.Create()
	log.ZapLog.With(zap.Any("生成记录g:", g)).Info("NewGameInst")

	if err != nil {
		log.ZapLog.With(zap.Any("生成记录err:", err)).Info("NewGameInst")
		return nil, err
	}

	inst := &GameInst{
		GameSid:   g.GameSid,
		startTime: g.Starttime,
		EnterNum:  0,
		BetNum:    0,
		c:         NewReady(),
		ss:        make(map[enum.State]ifstate),
		isEnd:     false,
	}

	hashRocketGame.Gi = inst
	return inst, nil
}

// SetEnd 设置获取本局游戏是否结束
func (i *GameInst) SetEnd() bool {
	for _, s := range i.ss {
		if s.GetStateDesc() == "end" {
			i.isEnd = true
			break
		}
	}

	i.isEnd = false
	return i.isEnd
}

func (i *GameInst) ChangeState() error {
	return nil
}
