package state

import (
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/supernet/common/net/gstream/pb"
	"github.com/supernet/game/hashrocket/enum"
	hrpb "github.com/supernet/game/hashrocket/internal/gstream/pb"
	"github.com/supernet/game/hashrocket/pkg/log"
)

type ready struct {
	state
}

func NewReady() ready {
	return ready{
		state: state{sd: "ready"},
	}
}

// CreateGame 准备阶段的逻辑
func (r ready) process() error {
	return nil
}

func (r ready) SetNextState() ifstate {
	return NewBetting()
}

func (r ready) GetStateDesc() string {
	return "ready"
}

func (r ready) GetState() enum.State {
	return enum.State(enum.READY)
}

func (r ready) init() (response []pb.StreamResponseData) {
	log.ZapLog.With(zap.Any("hashRocketGame.Gi:", hashRocketGame.Gi), zap.Any("hashRocketGame.Uids:", hashRocketGame.Uids)).Info("init")

	if len(hashRocketGame.Uids) == 0 {
		return
	}

	if hashRocketGame.Gi == nil || hashRocketGame.Gi.GameSid == "" {
		log.ZapLog.With(zap.Any("ready init:", "系统错误,没有初始化的游戏在运行")).Info("init")
		return
	}

	data := hrpb.GameEnterReplyToc{
		GameSid:   hashRocketGame.Gi.GameSid,
		StartTime: hashRocketGame.Gi.startTime.Unix(),
		EnterTime: time.Now().Unix(),
	}

	pd, _ := proto.Marshal(&data)
	// 发给所有的玩这个游戏的玩家
	sendCMsg := pb.StreamResponseData{
		ClientId: "",
		BAllUser: false,
		Uids:     hashRocketGame.Uids,
		Msg:      uint32(enum.CMD_READY),
		Data:     pd,
	}

	response = append(response, sendCMsg)
	return response
}

func (r ready) handleMsg(request *pb.StreamRequestData) (response []pb.StreamResponseData) {
	//获取
	if hashRocketGame.Gi == nil || hashRocketGame.Gi.GameSid == "" {
		log.ZapLog.With(zap.Any("ready init:", "系统错误,没有初始化的游戏在运行")).Error("init")
		return
	}

	date := hrpb.GameEnterReplyToc{
		GameSid:   hashRocketGame.Gi.GameSid,
		StartTime: hashRocketGame.Gi.startTime.Unix(),
		EnterTime: time.Now().Unix(),
	}

	log.ZapLog.With(zap.Any("handleMsg request.ClientId:", request.ClientId)).Info("init")
	data, _ := proto.Marshal(&date)
	sendCMsg := pb.StreamResponseData{
		ClientId: request.ClientId,
		BAllUser: false,
		Uids:     nil,
		Msg:      uint32(enum.CMD_ENTER_GAME),
		Data:     data,
	}
	userSid := strings.Split(request.ClientId, "_")[0]
	hashRocketGame.Uids = append(hashRocketGame.Uids, userSid)

	response = append(response, sendCMsg)
	return response
}
