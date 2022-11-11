package main

import (
	"flag"
	"fmt"

	"github.com/hewenyu/gozero-nacos-demo/common"
	"github.com/hewenyu/gozero-nacos-demo/rpc/internal/config"
	"github.com/hewenyu/gozero-nacos-demo/rpc/internal/server"
	"github.com/hewenyu/gozero-nacos-demo/rpc/internal/svc"
	"github.com/hewenyu/gozero-nacos-demo/rpc/pb"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/demo.yaml", "the nacos config file")

func main() {
	flag.Parse()

	var c config.Config
	// conf.MustLoad(*configFile, &c)

	common.MustRegister(common.MustLoad(*configFile, &c), &c.RpcServerConf)
	ctx := svc.NewServiceContext(&c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterDemoServer(grpcServer, server.NewDemoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
