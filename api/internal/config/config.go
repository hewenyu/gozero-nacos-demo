package config

import "github.com/zeromicro/go-zero/rest"

// Config 配置文件
type Config struct {
	rest.RestConf
	DemoRPC string // RPC 服务名称
}
