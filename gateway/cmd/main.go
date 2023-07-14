package main

import (
	"fmt"
	"github.com/supernet/gateway/internal/httpServer"
	"github.com/supernet/gateway/internal/net"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/supernet/gateway/pkg/log"

	_ "net/http/pprof"

	"github.com/supernet/gateway/pkg/mysql"
	"github.com/supernet/gateway/pkg/redislib"
	"github.com/supernet/gateway/pkg/viper"
)

func main() {
	// 初始化配置文件
	viper.InitVp()

	// 初始化日志文件
	log.ZapLog = log.InitLogger()
	log.ZapLog.With().Info("网关开始运行......")

	// 初始化redis
	redislib.Sclient()

	// 初始化数据库 获取 mysql.M()  mysql.S()
	MasterDB := mysql.MasterInit()
	defer MasterDB.Close()
	Slave1DB := mysql.Slave1Init()
	defer Slave1DB.Close()

	net.Init()
	go net.CManager.Run()
	defer net.CManager.Close()

	//wg := sync.WaitGroup{}
	//wg.Add(1)
	// 初始化grpc 连接池
	net.InitGrpcPools()
	net.ActiveGrpc()

	// http服务器启动
	go httpServer.Run()

	//信号捕捉 处理
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, os.Interrupt,
		syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.ZapLog.With(zap.Any("Signal", s)).Info("Program Exit...")
				GracefullExit()
			//避免网络断开系统错误,进程挂掉
			case syscall.SIGPIPE, os.Interrupt:
			default:
				log.ZapLog.With(zap.Any("Signal", s)).Info("other signal")
			}
		}
	}()

	go func() {
		runtime.SetMutexProfileFraction(1)
		runtime.SetBlockProfileRate(1)
		http.ListenAndServe("127.0.0.1:6060", nil)
	}()

	http.HandleFunc("/", net.WsHandler)

	port := viper.Vp.GetInt("ser.gateway.port")
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
}

// GracefullExit todo  等待收到的消息处理完毕 关闭客户端CManager 释放资源
func GracefullExit() {
	os.Exit(0)
}
