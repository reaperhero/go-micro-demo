package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go-micro-demo/14-micro-tools/proto"
	"go-micro-demo/14-micro-tools/usecase"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	myService := micro.NewService(
		micro.Name("api.mudu.com.app"),
		micro.Address(":8011"),
		micro.Registry(etcdReg),
	)

	proto.RegisterUsecaseServiceHandler(myService.Server(), new(usecase.ServiceImpl))

	myService.Init()

	err := myService.Run()
	if err != nil {
		fmt.Println(err)
	}
}
