package svc

import "github.com/hewenyu/gozero-nacos-demo/rpc/internal/config"

// ServiceContext 初始化
type ServiceContext struct {
	Config config.Config
}

// NewServiceContext 初始化
func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		Config: *c,
	}
}
