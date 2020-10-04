package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"go-micro-demo/13-micro-hystrix-threshold/web/models"
	"go-micro-demo/13-micro-hystrix-threshold/web/proto"
	"strconv"
)

// 通用降级方法
func defaultData(rsp interface{}) {
	switch t := rsp.(type) {
	case *models.ProdListResponse:
		defaultProds(rsp)
	case *models.ProdDetailResponse:
		t.Data = &models.ProdModel{ProdId: 10, ProdName: "降级商品"}
	}
}

func defaultProds(rsp interface{}) (*proto.ProdListResponse, error) {
	model := make([]*proto.ProdModel, 0)
	var i int32
	for i = 0; i < 3; i++ {
		prod := &proto.ProdModel{ProdId: 20 + i, ProdName: "prodName" + strconv.Itoa(20+int(i))}
		model = append(model, prod)
	}
	res := rsp.(*proto.ProdListResponse)
	res.Data = model
	return res, nil
}

// ProdsWrapper 商品装饰器
type ProdsWrapper struct {
	client.Client
}

// Call 调用方法
func (p *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	configA := hystrix.CommandConfig{
		Timeout:                1000, // 超时时间 单位毫秒
		RequestVolumeThreshold: 5,    // 请求数量
		ErrorPercentThreshold:  50,   // 错误百分比
		SleepWindow:            5000, // 尝试正常请求时间 单位毫秒 默认为5秒
	}
	hystrix.ConfigureCommand(cmdName, configA)
	return hystrix.Do(cmdName, func() error {
		return p.Client.Call(ctx, req, rsp)
	}, func(e error) error {
		defaultData(rsp)
		return nil
	})
	return p.Client.Call(ctx, req, rsp)
}

// NewProdsWrapper 初始化一个商品装饰器
func NewProdsWrapper(c client.Client) client.Client {
	return &ProdsWrapper{c}
}
