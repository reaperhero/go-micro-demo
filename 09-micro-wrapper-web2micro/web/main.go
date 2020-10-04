package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"go-micro-demo/08-micro-middlewares-web2micro/web/handler/http"
	"go-micro-demo/08-micro-middlewares-web2micro/web/proto"
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

func newLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	myService := micro.NewService(
		micro.Name("ProdService.client"),
		micro.WrapClient(newLogWrapper),
	)
	prodService := proto.NewProdService("ProdService", myService.Client()) // etcd里面查询ProdService服务，proto文件内容和rpc端需要一致

	service := web.NewService(
		web.Name("ProdService.client"),
		web.Address(":9000"),
		web.Handler(http.InitRouter(prodService)), // 把rpc客户端接口传入
		web.Registry(etcdReg),
	)

	service.Init()
	service.Run()
}
