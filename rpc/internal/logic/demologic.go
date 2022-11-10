package logic

import (
	"context"

	"github.com/hewenyu/gozero-nacos-demo/rpc/internal/svc"
	"github.com/hewenyu/gozero-nacos-demo/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

// DemoLogic 初始化
type DemoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewDemoLogic 初始化
func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// NewDemoLogic 初始化
func (l *DemoLogic) Demo(in *pb.DemoRequest) (*pb.DemoResponse, error) {

	return &pb.DemoResponse{
		Result: true,
	}, nil
}
