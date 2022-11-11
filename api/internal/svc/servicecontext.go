package svc

import (
	"github.com/hewenyu/gozero-nacos-demo/api/internal/config"
	"github.com/hewenyu/gozero-nacos-demo/common"
	"github.com/hewenyu/gozero-nacos-demo/rpc/demo"
)

// ServiceContext 初始化
type ServiceContext struct {
	Config  config.Config
	DemoRPC demo.Demo
}

// NewServiceContext 初始化
func NewServiceContext(c *config.Config, nc *common.Nacos) *ServiceContext {
	return &ServiceContext{
		Config:  *c,
		DemoRPC: demo.NewDemo(nc.NewZrpcClient(c.DemoRPC, c.Name)),
	}
}
