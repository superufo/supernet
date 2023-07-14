package routers

import (
	"github.com/supernet/gateway/pkg/log"
	"github.com/supernet/gateway/pkg/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var centerGrpc *grpc.ClientConn

func GetCenterGrpc() (*grpc.ClientConn, error) {
	if centerGrpc == nil {
		v := viper.Vp.GetStringSlice("center.addr")[0]
		url := string([]byte(v)[7:])
		conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.ZapLog.With(zap.Error(err)).Error("grpc dial error")
			return nil, err
		}
		centerGrpc = conn
		return centerGrpc, nil
	}
	return centerGrpc, nil
}
