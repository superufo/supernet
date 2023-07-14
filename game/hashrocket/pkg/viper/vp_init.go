package viper

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	Vp = viper.New()
)

func InitVp() {
	Vp.SetConfigName("config/config")
	Vp.AddConfigPath(".") // 添加搜索路径
	err := Vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败: %s \n", err))
	}
}
