package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro-demo/13-micro-hystrix-threshold/web/proto"
	"log"
)

// GetProdList 显示商品列表
func GetProdList(c *gin.Context) {
	prodService := c.Keys["prodservice"].(proto.ProdService)
	var prodReq proto.ProdRequest
	err := c.Bind(&prodReq)
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error()})
	} else {
		prodRes, _ := prodService.GetProdList(context.Background(), &prodReq)

		c.JSON(200, gin.H{"data": prodRes.Data})
	}

}

// GetProdDetail 显示商品详情
func GetProdDetail(c *gin.Context) {
	var prodReq proto.ProdRequest
	if err := c.BindUri(&prodReq); err != nil {
		log.Println(err)
		return
	}
	prodService := c.Keys["prodservice"].(proto.ProdService)
	res, _ := prodService.GetProdDetail(context.Background(), &prodReq)

	c.JSON(200, gin.H{"data": res.Data})
}
