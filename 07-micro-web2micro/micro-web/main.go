package main

import (
	"context"
	"go-micro-demo/07-micro-web2micro/micro-web/models"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	r := gin.Default()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	service := web.NewService(
		web.Name("testService.micro-rpc"),
		web.Address(":9000"),
		web.Handler(r),
		web.Registry(etcdReg),
	)

	myService := micro.NewService(micro.Name("tetsService.micro-rpc"))
	prodService := models.NewProdService("ProdService", myService.Client())
	v1Group := r.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", func(c *gin.Context) {
			var prodReq models.ProdRequest
			err := c.Bind(&prodReq)
			if err != nil {
				c.JSON(500, gin.H{
					"status": err.Error()})
			} else {
				prodRes, _ := prodService.GetProdList(context.Background(), &prodReq)
				c.JSON(200, gin.H{
					"data": prodRes.Data,
				})
			}
		})
	}

	service.Init()
	service.Run()
}

// curl --location --request POST 'localhost:9000/v1/prods' \
// --header 'Content-Type: application/json' \
// --data-raw '{"size":10}'
