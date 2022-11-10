package logic

import (
	"context"

	"github.com/hewenyu/gozero-nacos-demo/rpc/internal/svc"
	"github.com/hewenyu/gozero-nacos-demo/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DemoLogic) Demo(in *pb.DemoRequest) (*pb.DemoResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.DemoResponse{}, nil
}
