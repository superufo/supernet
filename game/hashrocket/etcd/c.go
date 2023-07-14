package etcd

import (
	"context"
	"fmt"
	"go.uber.org/zap"

	"github.com/supernet/common/discover/kit/sd/etcdv3"

	"github.com/supernet/game/hashrocket/config"
)

type client struct {
	cfg *config.ServerCfg
	log *zap.Logger

	c3 etcdv3.Client
}

func NewClient(cfg *config.ServerCfg, log *zap.Logger) (*client, error) {
	c3, err := etcdv3.NewClient(context.Background(), cfg.GetEtcdServer(), cfg.GetOption())
	if err != nil {
		log.With(zap.Any("err", err)).Error("NewClient")
		return nil, err
	}

	c := &client{
		cfg,
		log,
		c3,
	}
	return c, err
}

func (c *client) Reg() {
	instance := fmt.Sprintf("%s://%s:%d", c.cfg.GetProtocol(), c.cfg.GetIp(), c.cfg.GetPort())

	c.log.With(zap.Any("instance", instance)).Info("Reg")
	registrar := etcdv3.NewRegistrar(c.c3, etcdv3.Service{
		Key:   c.cfg.GetRegKey(),
		Value: instance,
	}, c.log)

	// Register our instance.
	registrar.Register()
	defer registrar.Deregister()
}

func (c *client) Discover() ([]string, error) {
	v, err := c.c3.GetEntries(c.cfg.GetRegKey())
	c.log.With(zap.Any("regKey", v)).Info("discover")

	return v, err
}
