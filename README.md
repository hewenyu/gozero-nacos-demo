# gozero-nacos-demo
go zero 使用nacos 2.x作为注册中心和配置中心 demo



## pb 文件

* [文件](./rpc/pb/demo.proto)
```bash
# 依赖生成
goctl rpc protoc rpc/pb/*.proto --go_out=rpc/ --go-grpc_out=rpc/  --zrpc_out=rpc/
sed -i  's/,omitempty//g' rpc/pb/*.pb.go
``