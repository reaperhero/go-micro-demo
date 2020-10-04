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
	// ProdId 商品Id
	ProdID int
	// ProdName 商品名称
	ProdName string
}

// NewProd 新增产品
func NewProd(id int, name string) *ProdModel {
	return &ProdModel{ProdID: id, ProdName: name}
}

// NewProdList 产品列表
func NewProdList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "Prod"+strconv.Itoa(100+i)))
	}
	return ret
}

// 调用函数返回json数据
func main() {

	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))

	r := gin.Default()
	// 路由分组
	v1Group := r.Group("/v1")
	{
		v1Group.Handle("GET", "/prods", func(c *gin.Context) {
			c.JSON(http.StatusOK, NewProdList(2))
		})
	}

	service := web.NewService(
		web.Name("ProdSrv"),
		web.Address(":8000"),
		web.Handler(r),
		web.Registry(etcdReg),
	)
	service.Run()
}

// /micro/registry/ProdSrv/46d1bed2-4063-43a1-9b4e-d81bff35840e
// {"name":"ProdSrv","version":"latest","metadata":null,"endpoints":null,"nodes":[{"id":"46d1bed2-4063-43a1-9b4e-d81bff35840e","address":"192.168.1.100:8000","metadata":null}]}