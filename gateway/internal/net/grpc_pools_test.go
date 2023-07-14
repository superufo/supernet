package net

import (
	"fmt"
	"github.com/supernet/gateway/pkg/log"
	"testing"

	"github.com/spf13/viper"
)

var (
	Vp = viper.New()
)

func TestGet(t *testing.T) {
	// 初始化配置文件
	Vp.SetConfigName("config/config")
	Vp.AddConfigPath("E:\\go_project\\game-server\\src\\github.com/supernet/gateway") // 添加搜索路径
	err := Vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败: %s \n", err))
	}

	// 初始化日志文件
	log.ZapLog = log.InitLogger()

	gps := NewGrpcPools("127.0.0.1:18089", 8)

	for i := 0; i < 1000; i++ {
		gps.Get("1111233244")
		gps.Get("111委任为他的功夫")
	}

	gps.Get("2234434344")
}

// BenchmarkAll也受TestMain限制
func BenchmarkGet(b *testing.B) {
	// 初始化配置文件
	Vp.SetConfigName("config/config")
	Vp.AddConfigPath("E:\\go_project\\game-server\\src\\github.com/supernet/gateway") // 添加搜索路径
	err := Vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败: %s \n", err))
	}

	// 初始化日志文件
	log.ZapLog = log.InitLogger()

	gps := NewGrpcPools("127.0.0.1:18089", 8)

	for n := 0; n < b.N; n++ {
		gps.Get("1111233244")
	}
}
