package main

import (
	"fmt"
	"github.com/supernet/game/hashrocket/internal/state"
	"go.uber.org/zap"
	"net/http"
	_ "net/http/pprof"
	"sync"

	"github.com/supernet/common/net/gstream/server"
	"github.com/supernet/game/hashrocket/config"

	"github.com/supernet/game/hashrocket/etcd"
	"github.com/supernet/game/hashrocket/pkg/log"
	"github.com/supernet/game/hashrocket/pkg/mysql"
	"github.com/supernet/game/hashrocket/pkg/redislib"
	"github.com/supernet/game/hashrocket/pkg/viper"
)

func main() {
	//http://127.0.0.1:6061/debug/pprof/
	go func() {
		err := http.ListenAndServe(":6061", nil) //设置监听的端口
		if err != nil {
			fmt.Printf("ListenAndServe: %s", err)
		}
	}()

	// 初始化配置文件
	viper.InitVp()

	// 初始化日志文件
	log.ZapLog = log.InitLogger()

	// 初始化redis
	redislib.Sclient()

	// 初始化数据库 获取 mysql.M()  mysql.S()
	MasterDB := mysql.MasterInit()
	defer MasterDB.Close()
	Slave1DB := mysql.Slave1Init()
	defer Slave1DB.Close()

	log.ZapLog.Info("哈希火箭服务器开始运行........")

	// 獲取配置信息
	scfg := config.NewServerCfg()

	/******* 服务注册 start*******/
	if c, err := etcd.NewClient(scfg, log.ZapLog); err == nil {
		c.Reg()
	}
	/******* 服务注册 start*******/

	grpcAddr := fmt.Sprintf("%s:%d", scfg.GetIp(), scfg.GetPort())
	var game *state.HashRocketGame
	once := sync.Once{}
	once.Do(func() {
		game = state.NewHashRocketGame()
	})

	log.ZapLog.With(zap.Any("", game)).Info("main........")

	//grpc流服务器运行游戏 游戏实例 日志 grpc地址
	server.RunGame(game, log.ZapLog, grpcAddr)
}
