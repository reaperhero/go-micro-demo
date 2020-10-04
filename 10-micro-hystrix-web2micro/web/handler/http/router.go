package http

import (
	"github.com/gin-gonic/gin"
	"go-micro-demo/10-micro-hystrix-web2micro/web/handler/middlewares"
	"go-micro-demo/10-micro-hystrix-web2micro/web/proto"
)

// InitRouter 路由
func InitRouter(prodService proto.ProdService) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.InitMiddleware(prodService))
	v1Group := r.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", GetProdList)
	}
	return r
}
