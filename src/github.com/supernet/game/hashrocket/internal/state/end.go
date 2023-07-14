package state

import (
	"github.com/supernet/common/net/gstream/pb"
	"github.com/supernet/game/hashrocket/enum"
)

type end struct {
	state
}

func NewEnd() end {
	return end{
		state: state{sd: "end"},
	}
}

// SetNextState 如果到結束了，状态为nil
func (r end) SetNextState() ifstate {
	return ready{}
}

func (r end) GetState() enum.State {
	return enum.State(enum.READY)
}

func (r end) GetStateDesc() string {
	return "end"
}

func (r end) init() []pb.StreamResponseData {
	return nil
}

func (r end) handleMsg(*pb.StreamRequestData) []pb.StreamResponseData {
	return nil
}
