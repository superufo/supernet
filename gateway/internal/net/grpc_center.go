package net

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/supernet/gateway/internal/httpServer/routers"
	"github.com/supernet/gateway/internal/model/entity"
	"github.com/supernet/gateway/internal/net/center/pb"
	"github.com/supernet/gateway/pkg/log"
)

// GetUserInfo grpc client 调用center 的用户信息
func GetUserInfo(sid string) (user *entity.Users, userInfo *entity.UserInfo, err error) {
	var response *pb.UserInfoResponse
	if centerConn, err := routers.GetCenterGrpc(); err != nil {
		log.ZapLog.With().Error("连接center server error")
	} else {
		storageClient := pb.NewStorageClient(centerConn)

		request := pb.UserRequest{
			Uid: sid,
		}

		response, err = storageClient.GetUserInfoFromDb(context.Background(), &request)
		if response.User != nil {
			copier.Copy(user, response.User)
		}
		if response.UserInfo != nil {
			copier.Copy(userInfo, response.UserInfo)
		}
	}

	return user, userInfo, err
}
