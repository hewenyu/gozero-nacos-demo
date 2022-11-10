# gozero-nacos-demo
go zero 使用nacos 2.x作为注册中心和配置中心 demo




## API 文件


```bash
# 使用命令生成api的代码
goctl api go -api ./api/desc/demo.api -dir ./api
```


```bash
# 进入到api的目录
cd api
# 使用命令生成api的代码
goctl api plugin -plugin goctl-swagger="swagger -filename demo.json" -api ./api/desc/demo.api -dir .
```

## pb 文件

* [文件](./rpc/pb/demo.proto)
```bash
# 依赖生成
goctl rpc protoc rpc/pb/*.proto --go_out=rpc/ --go-grpc_out=rpc/  --zrpc_out=rpc/
sed -i  's/,omitempty//g' rpc/pb/*.pb.go
```