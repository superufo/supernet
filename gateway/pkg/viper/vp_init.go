package viper

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var (
	Vp = viper.New()

	PsCfg = make([]ProxyCfg, 0)
)

func InitVp() {
	Vp.SetConfigName("config/config")
	Vp.AddConfigPath("../")
	Vp.AddConfigPath(".") // 添加搜索路径
	err := Vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件读取失败: %s \n", err))
	}

	proxys := Vp.Get("proxy")
	for i, p := range proxys.(map[string]interface{}) {
		fmt.Printf("i:%s proxy:%+v \n", i, p)

		proxy := ProxyCfg{}
		mapstructure.Decode(p, &proxy)
		PsCfg = append(PsCfg, proxy)
	}
	//fmt.Printf("PsCfg：%+v", PsCfg)
}
