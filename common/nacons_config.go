package common

import (
	"fmt"
	"log"
	"sync"

	"github.com/hewenyu/zero-contrib/zrpc/registry/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	zeroConf "github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

var (
	configClient config_client.IConfigClient
	nacosOnce    sync.Once
)

// Nacos 配置文件
type Nacos struct {
	Addr        string
	Port        uint64
	GrpcPort    uint64
	Group       string
	DataID      string
	NamespaceID string
	Username    string
	Password    string
	ExtDataIDs  []string `json:",optional"`
}

// InitConfigClient 初始化客户端
func (conf *Nacos) InitConfigClient() (err error) {
	nacosOnce.Do(func() {

		sc := []constant.ServerConfig{
			*constant.NewServerConfig(conf.Addr, conf.Port, constant.WithContextPath("/nacos"), constant.WithGrpcPort(conf.GrpcPort)),
		}

		// create ClientConfig
		cc := *constant.NewClientConfig(
			constant.WithNamespaceId(conf.NamespaceID),
			constant.WithTimeoutMs(5000),
			constant.WithNotLoadCacheAtStart(true),
			constant.WithLogDir("/tmp/nacos/log"),
			constant.WithCacheDir("/tmp/nacos/cache"),
			constant.WithUsername(conf.Username),
			constant.WithPassword(conf.Password),
			constant.WithLogLevel("info"),
			// constant.WithLogLevel("debug"),
		)

		configClient, err = clients.NewConfigClient(
			vo.NacosClientParam{
				ClientConfig:  &cc,
				ServerConfigs: sc,
			},
		)
		if err != nil {
			panic(err)
		}
	})
	return
}

// GetConfig 获取配置文件
func (conf *Nacos) GetConfig() (string, error) {

	mainConfig, err := configClient.GetConfig(vo.ConfigParam{DataId: conf.DataID, Group: conf.Group})
	if err != nil {
		return "", err
	}
	// logx.Debug(mainConfig)
	// return mainConfig, nil

	if len(conf.ExtDataIDs) == 0 {
		return mainConfig, nil
	}

	var configMap = make(map[interface{}]interface{})

	mainMap, err := UnmarshalYamlToMap(mainConfig)
	if err != nil {
		return "", err
	}

	var extMap = make(map[interface{}]interface{})
	for k := range conf.ExtDataIDs {
		extConfig, errMsg := configClient.GetConfig(vo.ConfigParam{DataId: conf.ExtDataIDs[k], Group: conf.Group})
		if err != nil {
			return "", errMsg
		}

		tmpExtMap, errMsg := UnmarshalYamlToMap(extConfig)
		if err != nil {
			return "", errMsg
		}

		extMap = MergeMap(extMap, tmpExtMap)
	}

	configMap = MergeMap(configMap, extMap)
	configMap = MergeMap(configMap, mainMap)

	yamlString, err := MarshalObjectToYamlString(configMap)
	if err != nil {
		return "", err
	}

	return yamlString, nil
}

// Listen 监听
func (conf *Nacos) Listen(onChange func(string, string, string, string)) error {
	return configClient.ListenConfig(vo.ConfigParam{
		DataId:   conf.DataID,
		Group:    conf.Group,
		OnChange: onChange,
	})
}

// NewZrpcClient RPC client
func (conf *Nacos) NewZrpcClient(serverName, clientName string) zrpc.Client {
	var target = fmt.Sprintf("nacos://%s:%s@%s:%d/%s?timeout=%s&namespace_id=%s&group_name=%s&app_name=%s&grpc=%d", conf.Username, conf.Password, conf.Addr, conf.Port, serverName, "10s", conf.NamespaceID, conf.Group, clientName, conf.GrpcPort)
	return zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: target,
	})
}

// MustLoad 配置获取
func MustLoad(nacosConfigFilePath string, v interface{}) *Nacos {
	var (
		err    error
		config string
	)

	var nacosConfig Nacos
	zeroConf.MustLoad(nacosConfigFilePath, &nacosConfig, zeroConf.UseEnv())
	err = nacosConfig.InitConfigClient()
	if err != nil {
		log.Fatalf("init config client error: %v", err)
	}

	config, err = nacosConfig.GetConfig()
	if err != nil {
		log.Fatalf("get config error: %v", err)
	}

	err = zeroConf.LoadConfigFromYamlBytes([]byte(config), v)
	if err != nil {
		log.Fatalf("load config error: %v", err)
	}
	return &nacosConfig
}

// MustRegister 注册
func MustRegister(nacosConfig *Nacos, rpcConfig *zrpc.RpcServerConf) {

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(nacosConfig.Addr, nacosConfig.Port, constant.WithContextPath("/nacos"), constant.WithGrpcPort(nacosConfig.GrpcPort)),
	}

	// create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(nacosConfig.NamespaceID),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithUsername(nacosConfig.Username),
		constant.WithPassword(nacosConfig.Password),
		constant.WithLogLevel("info"),
	)

	opts := nacos.NewNacosConfig(rpcConfig.Name, rpcConfig.ListenOn, sc, &cc, nacos.WithGroup(nacosConfig.Group))
	err := nacos.RegisterService(opts)
	if err != nil {
		log.Fatalf("register service failed: %s", err)
	}
}
