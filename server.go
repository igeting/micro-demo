package main

import (
	"github.com/asim/go-micro/v3"
	"micro-demo/handler"
	"micro-demo/proto/pb"
)

func main() {
	service := micro.NewService(
		micro.Name("test-service"),
	)
	service.Init()

	pb.RegisterHelloWorldHandler(service.Server(), new(handler.HelloWorld))

	service.Run()
}
