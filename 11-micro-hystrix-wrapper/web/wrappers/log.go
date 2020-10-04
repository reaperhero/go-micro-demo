package wrappers

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/v2/client"
)

// 装饰器设计模式：对于原有的方法增强功能
type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}
