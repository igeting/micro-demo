package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"

	"micro-demo/handler"
	"micro-demo/proto/pb"
)

func main() {
	service := micro.NewService(
		micro.Name("test-service"),
		micro.Address(":8081"),
		micro.Registry(consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))),
	)

	ser := service.Server()
	ser.Init(server.Advertise("localhost:8081"))

	pb.RegisterHelloWorldHandler(ser, new(handler.HelloWorld))

	service.Init()
	service.Run()
}
