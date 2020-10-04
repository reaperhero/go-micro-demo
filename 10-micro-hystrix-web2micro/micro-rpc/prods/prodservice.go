package prods

import (
	"context"
	"go-micro-demo/10-micro-hystrix-web2micro/micro-rpc/models"
	"log"
	"strconv"
)

// ProdService 商品服务
type ProdService struct{}

func newProd(id int32, pname string) *models.ProdModel {
	return &models.ProdModel{ProdId: id, ProdName: pname}
}

// GetProdList 返回商品列表
func (*ProdService) GetProdList(ctx context.Context, in *models.ProdRequest, res *models.ProdListResponse) error {
	log.Println(in.Size)
	models := make([]*models.ProdModel, 0)
	var i int32
	for i = 0; i < in.Size; i++ {
		models = append(models, newProd(100+i, "prodName"+strconv.Itoa(100+int(i))))
	}
	res.Data = models
	return nil
}
