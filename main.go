package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
)

func main() {
	service := micro.NewService(
		micro.Name("test-service"),
		micro.Registry(
			consul.NewRegistry(
				registry.Addrs("127.0.0.1:8500"),
			),
		),
	)
	service.Init()
	service.Run()
}
