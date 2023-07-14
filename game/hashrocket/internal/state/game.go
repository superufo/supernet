package state

import (
	"context"
	"github.com/supernet/common/net/gstream/pb"
	"github.com/supernet/game/hashrocket/enum"
	"github.com/supernet/game/hashrocket/pkg/log"
	"go.uber.org/zap"
	"time"
)

var (
	hashRocketGame HashRocketGame
)

type HashRocketGame struct {
	Online int64     //當前在綫人數
	Gi     *GameInst // 當前玩的這局游戲

	Gis  []*GameInst // 内存100局游戏的数据
	Uids []string    // 连接此游戏的所有玩家
}

func NewHashRocketGame() *HashRocketGame {
	inst, _ := NewGameInst()
	return &HashRocketGame{
		Online: 12,
		Gi:     inst,
		Gis:    make([]*GameInst, 100, 100),
		Uids:   make([]string, 0),
	}
}

// Run 状态机  准备阶段(创建游戏) 下注阶段  运行阶段  结束阶段
// 根据游戏不同的状态 下发不同消息给客户端
func (g HashRocketGame) Run(request <-chan *pb.StreamRequestData, response chan<- *pb.StreamResponseData) {
	for {
		if g.Gi == nil {
			log.ZapLog.With(zap.Any("ready init:", "系统错误，没有初始化游戏")).Error("init")
			continue
		}

		if g.Gi.c == nil {
			log.ZapLog.With(zap.Any("ready init:", "系统错误，游戏状态初始化错误")).Error("init")
			continue
		}

		//log.ZapLog.With(zap.Any("game状态:", g.Gi.c.GetStateDesc())).Info("Run")
		// 不需要收到消息，主动广播给所有客户端的消息
		in := g.Gi.c.init()
		if len(in) > 0 {
			for _, msg := range in {
				response <- &msg
			}
		}

		// 收到消息回复给每个客户端的消息
		ctx, cancelFunc := context.WithCancel(context.Background())
		go func(ctx context.Context) {
			for {
				select {
				case r := <-request:
					log.ZapLog.With(zap.Any("收到数据协议号r.Msg:", r.Msg)).Info("Run")
					if r != nil {
						// 处理每个客户端的请求
						rl := g.Gi.c.handleMsg(r)
						for _, msg := range rl {
							log.ZapLog.With(zap.Any("处理每个客户端的请求:", "ok")).Info("Run")
							response <- &msg
						}
					} else {
						log.ZapLog.With(zap.Any("请求为空:", "ok")).Info("Run")
					}
				case <-ctx.Done():
					return
				}
			}
		}(ctx)

		// 这个状态阶段逻辑处理
		g.Gi.c.process()

		//适当延迟2秒 处理客户端的网络延迟的包
		time.Sleep(1 * time.Second)
		// 传递给ctx.Done结束
		cancelFunc()

		//
		g.Gi.c = g.Gi.c.SetNextState()
		if g.Gi.c.GetState() == enum.State(enum.END) {
			g.Gi.SetEnd()
			g.Gis = append(g.Gis, g.Gi)
			inst, _ := NewGameInst()
			g.Gi = inst
		}
	}
}
