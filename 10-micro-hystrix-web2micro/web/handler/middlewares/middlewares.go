package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-micro-demo/10-micro-hystrix-web2micro/web/proto"
)

// InitMiddleware 注入prodService中间件
func InitMiddleware(prodService proto.ProdService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["prodservice"] = prodService
	}
}
