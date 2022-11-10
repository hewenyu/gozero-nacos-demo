package logic

import (
	"context"

	"github.com/hewenyu/gozero-nacos-demo/api/internal/svc"
	"github.com/hewenyu/gozero-nacos-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// PingLogic 测试使用
type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewPingLogic 测试使用
func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Ping 测试使用
func (l *PingLogic) Ping() (resp *types.PingResponse, err error) {

	return &types.PingResponse{
		Code: 0,
	}, nil
}
