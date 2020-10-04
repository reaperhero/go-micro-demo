package http

import (
	"github.com/gin-gonic/gin"
	"go-micro-demo/13-micro-hystrix-threshold/web/handler/middlewares"
	"go-micro-demo/13-micro-hystrix-threshold/web/proto"
)

// InitRouter 路由
func InitRouter(prodService proto.ProdService) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.InitMiddleware(prodService), middlewares.ErrorMiddleware())
	v1Group := r.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", GetProdList)
		v1Group.Handle("GET", "/prods/:pid", GetProdDetail)
	}
	return r
}
