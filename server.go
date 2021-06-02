package main

import (
	"github.com/micro/go-micro/v2"
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
