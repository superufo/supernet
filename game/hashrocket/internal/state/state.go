package state

import (
	"context"
	"github.com/supernet/common/net/gstream/pb"
	"github.com/supernet/game/hashrocket/enum"
)

type state struct {
	ctx context.Context

	// 状态的描述
	sd string
}

func (s state) process() error {
	return nil
}

func (s state) SetNextState() ifstate {
	return state{}
}

func (s state) GetState() enum.State {
	return enum.State(enum.READY)
}

func (s state) GetStateDesc() string {
	return s.sd
}

func (s state) init() []pb.StreamResponseData {
	return nil
}

func (s state) handleMsg(*pb.StreamRequestData) []pb.StreamResponseData {
	return nil
}
