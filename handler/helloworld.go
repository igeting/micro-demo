package handler

import (
	"context"
	"micro-demo/proto/pb"
)

type HelloWorld struct {
}

func (HelloWorld) SayHello(ctx context.Context, req *pb.Request, res *pb.Response) error {
	res.Result = "hello:" + req.Name
	return nil
}
