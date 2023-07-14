package main

import (
	"fmt"
	"github.com/supernet/common/discover/kit"
	"github.com/supernet/common/log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"

	"context"

	. "github.com/supernet/common/discover/kit/sd/etcdv3"
)

func main() {
	var (
		etcdServer = "127.0.0.1:2379"    // in the change from v2 to v3, the schema is no longer necessary if connecting directly to an etcd v3 instance
		prefix     = "/services/foosvc/" // known at compile time
		instance   = "127.0.0.1:8080"    // taken from runtime or platform, somehow
		key        = prefix + instance   // should be globally unique

		value = "http://" + instance // based on our transport
		ctx   = context.Background()
	)

	options := ClientOptions{
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
	}

	// Build the client.
	client, err := NewClient(ctx, []string{etcdServer}, options)
	if err != nil {
		panic(err)
	}

	logger := log.NewLogger(
		log.SetPath("/log"),
		log.SetPrefix("test"),
		log.SetDebugFileSuffix("debug.log"),
		log.SetWarnFileSuffix("warn.log"),
		log.SetErrorFileSuffix("error.log"),
		log.SetInfoFileSuffix("info.log"),
		log.SetMaxAge(2),
		log.SetMaxBackups(30),
		log.SetMaxSize(10),
		log.SetDevelopment(true),
		log.SetLevel(zap.DebugLevel),
	)

	kit.SetKitLogger(logger)

	// Build the registrar.
	registrar := NewRegistrar(client, Service{
		Key:   key,
		Value: value,
	}, logger)

	// Register our instance.
	registrar.Register()

	// At the end of our service lifecycle, for example at the end of func main,
	// we should make sure to deregister ourselves. This is important! Don't
	// accidentally skip this step by invoking a log.Fatal or os.Exit in the
	// interim, which bypasses the defer stack.
	defer registrar.Deregister()

	v, _ := client.GetEntries(key)

	logger.With(zap.Any("v", v)).Info("")

	fmt.Print(v)

}
