package usecase

import (
	"context"
	"go-micro-demo/14-micro-tools/proto"
	"strconv"
)

type ServiceImpl struct {
}

func (*ServiceImpl) Call(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Data = "test" + strconv.Itoa(int(req.Id))
	return nil
}
