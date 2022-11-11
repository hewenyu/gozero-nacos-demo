# gozero-nacos-demo
go zero 使用nacos 2.x作为注册中心和配置中心 demo


```bash
# 使用前
go get -u github.com/hewenyu/zero-contrib/zrpc/registry/nacos
```

## API 文件


```bash
# 使用命令生成api的代码
goctl api go -api ./api/desc/demo.api -dir ./api
```


```bash
# 使用命令生成api的 swager 文件
goctl api plugin -plugin goctl-swagger="swagger -filename demo.json" -api ./api/desc/demo.api -dir .
```

## PB 文件

* [文件](./rpc/pb/demo.proto)
```bash
# 依赖生成
goctl rpc protoc rpc/pb/*.proto --go_out=rpc/ --go-grpc_out=rpc/  --zrpc_out=rpc/
sed -i  's/,omitempty//g' rpc/pb/*.pb.go
```

