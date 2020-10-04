package main

import (
	"github.com/micro/go-micro/v2"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"go-micro-demo/13-micro-hystrix-threshold/web/handler/http"
	"go-micro-demo/13-micro-hystrix-threshold/web/proto"
	"go-micro-demo/13-micro-hystrix-threshold/web/wrappers"
)

func main() {

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	myService := micro.NewService(
		micro.Name("ProdService.client"),
		micro.WrapClient(wrappers.NewLogWrapper),
		micro.WrapClient(wrappers.NewProdsWrapper),
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
