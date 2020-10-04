package http

import (
	"context"
	"go-micro-demo/08-micro-middlewares-web2micro/web/proto"

	"github.com/gin-gonic/gin"
)

// GetProdList 显示商品列表
func GetProdList(c *gin.Context) {
	prodService := c.Keys["prodservice"].(proto.ProdService) // 从middleware容器中获取，并断言
	var prodReq proto.ProdRequest
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

}
