package state

import (
	"github.com/supernet/common/net/gstream/pb"
	"github.com/supernet/game/hashrocket/enum"
)

type betting struct {
	state
}

func NewBetting() betting {
	return betting{
		state: state{sd: "betting"},
	}
}

func (t betting) SetNextState() ifstate {
	return running{}
}

func (t betting) GetStateDesc() string {
	return "betting"
}

func (t betting) GetState() enum.State {
	return enum.State(enum.BETING)
}

func (t betting) init() []pb.StreamResponseData {
	return nil
}

func (t betting) handleMsg(*pb.StreamRequestData) []pb.StreamResponseData {
	return nil
}
