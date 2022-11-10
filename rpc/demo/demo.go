// Code generated by goctl. DO NOT EDIT!
// Source: demo.proto

package demo

import (
	"context"

	"github.com/hewenyu/gozero-nacos-demo/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DemoRequest  = pb.DemoRequest
	DemoResponse = pb.DemoResponse

	Demo interface {
		Demo(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error)
	}

	defaultDemo struct {
		cli zrpc.Client
	}
)

func NewDemo(cli zrpc.Client) Demo {
	return &defaultDemo{
		cli: cli,
	}
}

func (m *defaultDemo) Demo(ctx context.Context, in *DemoRequest, opts ...grpc.CallOption) (*DemoResponse, error) {
	client := pb.NewDemoClient(m.cli.Conn())
	return client.Demo(ctx, in, opts...)
}