package main

import (
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"log"
)


func main() {
	// etcd连接句柄
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))

	// 获取服务
	getService, err := etcdReg.GetService("ProdSrv")
	if err != nil {
		log.Fatalf("get service failed, err:%v\n", err)
		return
	}

	next := selector.Random(getService) // 还有一种RoundRobin

	node, err := next()
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 打印服务节点信息
	log.Println(node.Address) // 打印节点信息  192.168.1.100:8000
}