package redislib

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/supernet/gateway/pkg/viper"
)

var (
	client *redis.Client
)

func Sclient() {
	ctx := context.Background()
	client = redis.NewClient(&redis.Options{
		Addr:         viper.Vp.GetString("redis.addr"),
		Password:     viper.Vp.GetString("redis.password"),
		DB:           viper.Vp.GetInt("redis.DB"),
		PoolSize:     viper.Vp.GetInt("redis.poolSize"),
		MinIdleConns: viper.Vp.GetInt("redis.minIdleConns"),
	})
	pong, err := client.Ping(ctx).Result()
	fmt.Println("初始化redis:", pong, err)
	// Output: PONG <nil>
}

func GetClient() (c *redis.Client) {
	return client
}
