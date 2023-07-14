package common

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"

	"github.com/supernet/gateway/internal/net/common/pb"
	"github.com/supernet/gateway/pkg/log"
)

type Service struct {
}

func NewServer() *Service {
	return &Service{}
}

func (s *Service) Call(context.Context, *pb.Request) (*pb.Response, error) {
	// 根据 Msg 头的信息,返回对应的ResponseData
	return nil, nil
}

func (s *Service) mustEmbedUnimplementedHallServer() {}

func Run() {
	var server pb.CommonServer
	sImpl := NewServer()

	server = sImpl

	// keepalive
	g := grpc.NewServer()

	// 2.注册逻辑到server中
	pb.RegisterCommonServer(g, server)

	instance := fmt.Sprintf("0.0.0.0:11190")
	log.ZapLog.With(zap.Any("addr", instance)).Info("Run")
	// 3.启动server
	lis, err := net.Listen("tcp", instance)
	if err != nil {
		panic("监听错误:" + err.Error())
	}

	err = g.Serve(lis)
	if err != nil {
		panic("启动错误:" + err.Error())
	}
}
