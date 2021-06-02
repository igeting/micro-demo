package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"log"
	"micro-demo/file/handler"
	"micro-demo/file/proto/pb"
)

func main() {
	service := micro.NewService(
		micro.Name("test-service"),
		micro.Address("localhost:8081"),
		micro.Registry(
			consul.NewRegistry(
				registry.Addrs("127.0.0.1:8500"),
			),
		),
	)

	err := pb.RegisterFileHandler(service.Server(), new(handler.FileManager))
	if err != nil {
		log.Fatalln(err)
	}

	service.Init()
	service.Run()
}
