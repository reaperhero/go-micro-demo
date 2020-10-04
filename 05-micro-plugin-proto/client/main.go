package main

import (
	"context"
	"go-micro-demo/05-micro-plugin-proto/client/models"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/etcd"
)

// 调用http api 引入protobuf生成请求响应模型
func callAPI(s selector.Selector) {
	myClient := http.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myClient.NewRequest("ProdSrv", "/v1/prods",models.ProdRequest{Size: 3})
	var rsp models.ProdListResponse
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(rsp.GetData()) // [ProdId:100 ProdName:"Prod100"  ProdId:101 ProdName:"Prod101"  ProdId:102 ProdName:"Prod102" ]
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
