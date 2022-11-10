// Code generated by goctl. DO NOT EDIT!
// Source: demo.proto

package server

import (
	"context"

	"github.com/hewenyu/gozero-nacos-demo/rpc/internal/logic"
	"github.com/hewenyu/gozero-nacos-demo/rpc/internal/svc"
	"github.com/hewenyu/gozero-nacos-demo/rpc/pb"
)

type DemoServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedDemoServer
}

func NewDemoServer(svcCtx *svc.ServiceContext) *DemoServer {
	return &DemoServer{
		svcCtx: svcCtx,
	}
}

func (s *DemoServer) Demo(ctx context.Context, in *pb.DemoRequest) (*pb.DemoResponse, error) {
	l := logic.NewDemoLogic(ctx, s.svcCtx)
	return l.Demo(in)
}