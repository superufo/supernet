package mongolib

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/supernet/gateway/pkg/log"
	viper2 "github.com/supernet/gateway/pkg/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MgDatabase struct {
	Mongo *mongo.Client
	Con   *mongo.Database
}

var MongoDB *MgDatabase

//初始化
func Init() {
	_con := SetConnect()
	MongoDB = &MgDatabase{
		Mongo: _con,
		Con:   _con.Database(viper2.Vp.GetString("mongod-master.dbname")),
	}
}

// 连接设置
func SetConnect() *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString("mongod-master.addr"), viper.GetString("mongod-master.port"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(20)) // 连接池
	if err != nil {
		log.ZapLog.Error(err.Error())
	}
	return client
}
