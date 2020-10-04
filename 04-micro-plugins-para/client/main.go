package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/etcd"
)

// 使用插件 调用http api 带参数调用
func callAPI(s selector.Selector) {
	myClient := http.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myClient.NewRequest("ProdSrv", "/v1/prods", map[string]interface{}{"size": 3})
	var rsp map[string]interface{}
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(rsp["data"]) // [map[ProdID:100 ProdName:Prod100] map[ProdID:101 ProdName:Prod101] map[ProdID:102 ProdName:Prod102]]
}

func main() {
	// etcd连接句柄
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))

	sel := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.RoundRobin),
	)
	callAPI(sel)

}
