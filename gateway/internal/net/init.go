package net

import (
	"github.com/supernet/gateway/pkg/viper"
	"time"
)

var (
	PoolsCollect = make(map[string]GrpcPools)
)

func InitGrpcPools() {
	for _, proxy := range viper.PsCfg {
		go func(p viper.ProxyCfg) {
			for _, url := range p.Addr {
				b := []byte(url)
				u := string(b[7:])

				PoolsCollect[u] = NewGrpcPools(u, p.Maxgrpc)
			}
		}(proxy)
	}
}

func ActiveGrpc() {
	ticker := time.NewTicker(10 * time.Second)

	go func(t *time.Ticker) {
		defer t.Stop()

		for {
			select {
			case <-t.C:
				//log.ZapLog.With().Info("激活.............")
				for url, p := range PoolsCollect {
					p.ActiveDeadLink(url)
				}
			}
		}
	}(ticker)
}
