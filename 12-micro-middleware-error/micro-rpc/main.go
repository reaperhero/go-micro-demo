package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go-micro-demo/12-micro-middleware-error/micro-rpc/proto"
	"go-micro-demo/12-micro-middleware-error/micro-rpc/usecase"
)

func main() {

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	app := micro.NewService(
		micro.Name("ProdService"),
		micro.Address(":8000"),
		micro.Registry(etcdReg),
	)

	app.Init()
	err := proto.RegisterProdServiceHandler(app.Server(), new(usecase.ProdService))
	if err != nil {
		panic(err)
	}
	app.Run()
}
