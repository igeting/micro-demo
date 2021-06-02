package main

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"log"
	"micro-demo/proto/pb"
)

func main() {
	service := micro.NewService(
		micro.Registry(consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))),
	)
	service.Init()

	cli := pb.NewHelloWorldService("test-service", service.Client())

	res, err := cli.SayHello(context.Background(), &pb.Request{
		Name: "lulu",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}
