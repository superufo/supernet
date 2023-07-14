package state

import (
	"github.com/supernet/common/net/gstream/pb"
	"github.com/supernet/game/hashrocket/enum"
)

type ifstate interface {
	SetNextState() ifstate
	GetState() enum.State

	// 初始化需要发送给客户端的信息
	init() []pb.StreamResponseData
	// 处理请求信息 回复给客户端的信息
	handleMsg(*pb.StreamRequestData) []pb.StreamResponseData
	GetStateDesc() string

	// 状态机的运行逻辑代码
	process() error
}
