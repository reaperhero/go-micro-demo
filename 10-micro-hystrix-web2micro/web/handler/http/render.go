package http

import (
	"context"
	"go-micro-demo/10-micro-hystrix-web2micro/web/proto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func newProd(id int32, pname string) *proto.ProdModel {
	return &proto.ProdModel{ProdId: id, ProdName: pname}
}

func defaultProds() (*proto.ProdListResponse, error) {
	model := make([]*proto.ProdModel, 0)
	var i int32
	for i = 0; i < 2; i++ {
		model = append(model, newProd(20+i, "defaultname"+strconv.Itoa(20+int(i))))
	}
	res := &proto.ProdListResponse{}
	res.Data = model
	return res, nil
}

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

		c.JSON(200, gin.H{"data": prodRes.Data})
	}

}
