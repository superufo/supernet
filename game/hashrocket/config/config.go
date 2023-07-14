package config

import (
	"github.com/supernet/game/hashrocket/pkg/viper"
	"google.golang.org/grpc"
	"math/rand"
	"time"

	"github.com/supernet/common/discover/kit/sd/etcdv3"
)

type ServerCfg struct {
	protocol   string
	ip         string
	port       int
	etcdServer []string
	regKey     string

	option etcdv3.ClientOptions
}

func NewServerCfg() *ServerCfg {
	protocol := viper.Vp.GetString("ser.protocol")
	ip := viper.Vp.GetString("ser.ip")
	port := viper.Vp.GetInt("ser.port")
	etcdServer := viper.Vp.GetStringSlice("ser.etcdServer")
	regKey := viper.Vp.GetString("ser.regKey")

	return &ServerCfg{
		protocol:   protocol,
		ip:         ip,
		port:       port,
		etcdServer: etcdServer,
		regKey:     regKey,
		option: etcdv3.ClientOptions{
			// Path to trusted ca file
			CACert: "",
			// Path to certificate
			Cert: "",
			// Path to private key
			Key: "",
			// Username if required
			Username: "",
			// Password if required
			Password: "",
			// If DialTimeout is 0, it defaults to 3s
			DialTimeout: time.Second * 3,
			// If DialKeepAlive is 0, it defaults to 3s
			DialKeepAlive: time.Second * 3,
			// If passing `grpc.WithBlock`, dial connection will block until success.
			DialOptions: []grpc.DialOption{grpc.WithBlock()},
		},
	}
}

func (f *ServerCfg) GetProtocol() string {
	return f.protocol
}

func (f *ServerCfg) GetIp() string {
	return f.ip
}

func (f *ServerCfg) GetPort() int {
	return f.port
}

func (f *ServerCfg) GetEtcdServer() []string {
	return f.etcdServer
}

func (f *ServerCfg) GetRandEtcdServer() string {
	n := rand.Intn(len(f.etcdServer))
	return f.etcdServer[n]
}

func (f *ServerCfg) GetRegKey() string {
	return f.regKey
}

func (f *ServerCfg) GetOption() etcdv3.ClientOptions {
	return f.option
}
