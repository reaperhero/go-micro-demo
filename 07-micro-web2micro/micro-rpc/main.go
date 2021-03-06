package main

import (
	"go-micro-demo/07-micro-web2micro/micro-rpc/models"
	"go-micro-demo/07-micro-web2micro/micro-rpc/prods"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
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
	err := models.RegisterProdServiceHandler(app.Server(), new(prods.ProdService))
	if err != nil {
		panic(err)
	}
	app.Run()
}
