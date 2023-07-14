package routers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/supernet/common/def"
	proto "github.com/supernet/gateway/internal/net/center/pb"
	"github.com/supernet/gateway/pkg/log"
	"go.uber.org/zap"
)

var admin = new(Admin)

const (
	UpdateConfig  = "update_config"  //更新配置
	ReduceBalance = "reduce_balance" //减钱
	AddBalance    = "add_balance"    //加钱
)

type AdminRequest struct {
	MainType  string `json:"msg_main_type"`
	SubType   string `json:"msg_sub_type"`
	SId       string `json:"s_id"`
	ChangGold int64  `json:"change_gold"`
}

type Admin struct {
}

func (a *Admin) Test(c *gin.Context) {
	c.JSON(200, Response{
		Code: 0,
		Msg:  "test ok",
		Data: nil,
	})
}

func (a *Admin) MsgHandler(c *gin.Context) {
	params := AdminRequest{}
	c.BindJSON(&params)
	switch params.MainType {
	case UpdateConfig:
		a.UpdateConfig(c, &params)
		return
	case ReduceBalance:
		a.ReduceBalance(c, &params)
		return
	case AddBalance:
		a.AddBalance(c, &params)
		return
	}
	c.JSON(200, Response{
		Code: 1,
		Msg:  ERROR_MESSAGE_TYPE,
		Data: nil,
	})
}

func (a *Admin) UpdateConfig(c *gin.Context, req *AdminRequest) {
	c.JSON(200, Response{
		Code: 0,
		Msg:  SUCCESS,
		Data: nil,
	})
}

func (a *Admin) ReduceBalance(c *gin.Context, req *AdminRequest) {
	if req.ChangGold <= 0 {
		c.JSON(200, Response{
			Code: 1,
			Msg:  INVALID_PARAMS,
			Data: nil,
		})
		return
	}
	center, err := GetCenterGrpc()
	if err != nil {
		c.JSON(200, Response{
			Code: 1,
			Msg:  ERROR_SERVER_GRPC,
			Data: nil,
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := proto.NewStorageClient(center)
	request := proto.ChangeBalanceReq{
		Uid:        req.SId,
		Gold:       req.ChangGold,
		ChangeType: uint32(def.CHANGE_ADMINREDUCE),
	}
	reply, err := client.ReduceBalance(ctx, &request)
	if err != nil {
		log.ZapLog.With(zap.Error(err)).Error("grpc dial result")
		c.JSON(200, Response{
			Code: 1,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, Response{
		Code: 0,
		Msg:  SUCCESS,
		Data: struct {
			BeforeGold int64 `json:"before_gold"`
			AfterGold  int64 `json:"after_gold"`
		}{
			BeforeGold: reply.GetBeforeGold(),
			AfterGold:  reply.GetAfterGold(),
		},
	})
}

func (a *Admin) AddBalance(c *gin.Context, req *AdminRequest) {
	if req.ChangGold <= 0 {
		c.JSON(200, Response{
			Code: 1,
			Msg:  INVALID_PARAMS,
			Data: nil,
		})
		return
	}
	center, err := GetCenterGrpc()
	if err != nil {
		c.JSON(200, Response{
			Code: 1,
			Msg:  ERROR_SERVER_GRPC,
			Data: nil,
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := proto.NewStorageClient(center)
	request := proto.ChangeBalanceReq{
		Uid:        req.SId,
		Gold:       req.ChangGold,
		ChangeType: uint32(def.CHANGE_ADMINADD),
	}
	reply, err := client.AddBalance(ctx, &request)
	if err != nil {
		log.ZapLog.With(zap.Error(err)).Error("grpc dial result")
		c.JSON(200, Response{
			Code: 1,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}
	c.JSON(200, Response{
		Code: 0,
		Msg:  SUCCESS,
		Data: struct {
			BeforeGold int64 `json:"before_gold"`
			AfterGold  int64 `json:"after_gold"`
		}{
			BeforeGold: reply.GetBeforeGold(),
			AfterGold:  reply.GetAfterGold(),
		},
	})
}
