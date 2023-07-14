package state

import (
	"github.com/supernet/common/net/gstream/pb"
	"github.com/supernet/game/hashrocket/enum"
)

type running struct {
	state
}

func NewRunning() running {
	return running{
		state: state{sd: "running"},
	}
}

func (r running) SetNextState() ifstate {
	return end{}
}

func (r running) GetState() enum.State {
	return enum.State(enum.READY)
}

func (r running) GetStateDesc() string {
	return "running"
}

func (r running) init() []pb.StreamResponseData {
	return nil
}

func (r running) handleMsg(*pb.StreamRequestData) []pb.StreamResponseData {
	return nil
}
