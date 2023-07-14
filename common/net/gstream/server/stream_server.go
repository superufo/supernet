package server

import (
	"github.com/supernet/common/net/gstream"
	pb2 "github.com/supernet/common/net/gstream/pb"
	Utils "github.com/supernet/common/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"

	"errors"
	"fmt"
	"io"
)

// var center  = storage.StorageServerImpl
var stream *streamServer

type streamServer struct {
	GrpcRecvClientData chan *pb2.StreamRequestData
	GrpcSendClientData chan *pb2.StreamResponseData

	Game *gstream.Game
	log  *zap.Logger
}

// NewStreamServer 一个客户端可以缓存 100 消息
func NewStreamServer(game *gstream.Game, log *zap.Logger) *streamServer {
	stream = &streamServer{
		make(chan *pb2.StreamRequestData, 100),
		make(chan *pb2.StreamResponseData, 100),
		game,
		log,
	}

	return stream
}

// PPStream log.ZapLog.With(zap.Any("err", err)).Error("收到网关数据错误")
func (gs *streamServer) PPStream(stream pb2.ForwardMsg_PPStreamServer) error {
	err := make(chan error)

	defer func() {
		if e := recover(); e != nil {
			gs.log.With(zap.Any("recover err", e.(error))).Info("PPStream")
			fmt.Print(e)
		}
	}()

	//once := sync.Once{}
	//once.Do(func() {
	//	gs.Game.Run(gs.GrpcRecvClientData, gs.GrpcSendClientData)
	//})
	if gs.Game != nil {
		go (*(gs.Game)).Run(gs.GrpcRecvClientData, gs.GrpcSendClientData)
	}

	go gs.response(stream, err)
	go gs.request(stream, err)

	return <-err
}

func (gs *streamServer) request(stream pb2.ForwardMsg_PPStreamServer, errCh chan error) {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			gs.log.With(zap.Any("PPStream recv io EOF", err)).Info("request")
			errCh <- err
			continue
		}

		if err != nil {
			gs.log.With(zap.Any("PPStream recv error", err)).Info("request")
			errCh <- err
			continue
		}

		info := fmt.Sprintf("收到网关数据:协议号=%+v,protobuf=%+v", msg.GetMsg(), Utils.ToHexString(msg.GetData()))
		gs.log.With(zap.Any("info:", info)).Info("request")
		gs.GrpcRecvClientData <- msg
	}
}

func (gs *streamServer) response(stream pb2.ForwardMsg_PPStreamServer, errCh chan error) {
	defer func() {
		if e := recover(); e != nil {
			gs.log.With(zap.Any("PPStream recv error", e.(error))).Info("stream response")
		}
	}()

	for {
		select {
		case sd := <-gs.GrpcSendClientData:
			//业务代码
			if err := stream.Send(sd); err != nil {
				err := errors.New("发给网关失败err:" + err.Error())
				gs.log.With(zap.Any("发给网关失败err", err)).Info("response")
				errCh <- err
			} else {
				gs.log.With(zap.Any("发给网关成功msg", sd.String())).Info("response")
			}
		}
	}
}

func RunGame(game gstream.Game, log *zap.Logger, grpcAddr string) {
	var server pb2.ForwardMsgServer
	sImpl := NewStreamServer(&game, log)

	server = sImpl

	g := grpc.NewServer()
	// 2.注册逻辑到server中
	pb2.RegisterForwardMsgServer(g, server)

	log.With(zap.Any("grpcAddr", grpcAddr)).Info("RunGame")
	// 3.启动server
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic("监听错误:" + err.Error())
	}

	err = g.Serve(lis)
	if err != nil {
		panic("启动错误:" + err.Error())
	} else {
		//log.With(zap.Any("运行游戏", sImpl.GrpcSendClientData)).Info("RunGame 运行游戏")
		//go sImpl.Game.Run(sImpl.GrpcRecvClientData, sImpl.GrpcSendClientData)
	}
}
