package main

import (
	"flag"
	"fmt"

	"github.com/hewenyu/gozero-nacos-demo/api/internal/config"
	"github.com/hewenyu/gozero-nacos-demo/api/internal/handler"
	"github.com/hewenyu/gozero-nacos-demo/api/internal/svc"
	"github.com/hewenyu/gozero-nacos-demo/common"
	_ "github.com/hewenyu/zero-contrib/zrpc/registry/nacos"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/demo-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	common.MustRegisterAPI(common.MustLoad(*configFile, &c), &c.RestConf)
	ctx := svc.NewServiceContext(&c, common.MustLoad(*configFile, &c))
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
