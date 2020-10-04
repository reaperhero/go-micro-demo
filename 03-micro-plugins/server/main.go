package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"net/http"
	"strconv"
)

// ProdModel 商品模型
type ProdModel struct {
	ProdID   int
	ProdName string
}

// NewProd 新增商品
func NewProd(id int, name string) *ProdModel {
	return &ProdModel{ProdID: id, ProdName: name}
}

// NewProdList 根据前端请求的size返回商品数量
func NewProdList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "Prod"+strconv.Itoa(100+i)))
	}
	return ret
}

// 商品服务
func main() {

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))

	r := gin.Default()
	// 路由分组
	v1Group := r.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": NewProdList(2),
			})
		})
	}

	service := web.NewService(
		web.Address(":8000"),
		web.Name("ProdSrv"),
		web.Handler(r),
		web.Registry(etcdReg),
	)

	// 通过命令行参数启动
	// --server_address 指定地址端口，或者环境变量$MICRO_SERVER_ADDRESS]
	// 运行2个服务
	// go run main.go prodModels.go --server_address  127.0.0.1:8000
	// go run main.go prodModels.go --server_address  127.0.0.1:8001
	//service.Init()
	service.Run()
}
