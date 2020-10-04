package usecase

import (
	"context"
	"go-micro-demo/12-micro-middleware-error/micro-rpc/proto"
	"log"
	"strconv"
	"time"
)

// ProdService 商品服务
type ProdService struct{}

func newProd(id int32, pname string) *proto.ProdModel {
	return &proto.ProdModel{ProdId: id, ProdName: pname}
}

// GetProdList 返回商品列表
func (*ProdService) GetProdList(ctx context.Context, in *proto.ProdRequest, res *proto.ProdListResponse) error {
	log.Println(in.Size)
	// 超时测试
	time.Sleep(time.Second * 3)
	proto := make([]*proto.ProdModel, 0)
	var i int32
	for i = 0; i < in.Size; i++ {
		proto = append(proto, newProd(100+i, "prodName"+strconv.Itoa(100+int(i))))
	}
	res.Data = proto
	return nil
}

// GetProdDetail 获取单个商品
func (*ProdService) GetProdDetail(ctx context.Context, in *proto.ProdRequest, res *proto.ProdDetailResponse) error {
	res.Data = newProd(in.ProdId, "测试商品")
	return nil
}
