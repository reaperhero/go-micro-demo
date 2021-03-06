package http

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"go-micro-demo/08-micro-middlewares-web2micro/web/proto"
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
		// 超时代码
		// 1.配置config
		configA := hystrix.CommandConfig{
			Timeout: 1000,
		}
		// 2.配置command
		hystrix.ConfigureCommand("getProds", configA)
		// 3.执行Do方法
		var prodRes *proto.ProdListResponse
		err := hystrix.Do("getProds", func() error {
			prodRes, err = prodService.GetProdList(context.Background(), &prodReq)
			return err
		}, nil)
		if err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
		} else {
			c.JSON(200, gin.H{"data": prodRes.Data})
		}
	}

}
