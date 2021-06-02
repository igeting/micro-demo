package main

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"log"
	"micro-demo/file/proto/pb"
	"os"
)

func main() {
	service := micro.NewService(
		micro.Registry(
			consul.NewRegistry(
				registry.Addrs("127.0.0.1:8500"),
			),
		),
	)
	service.Init()

	cli := pb.NewFileService("test-service", service.Client())

	path := "/home/c/logback.png"

	info, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
	}
	fi, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := cli.Upload(context.TODO(), &pb.FileRequest{
		Name: info.Name(),
		Size: info.Size(),
		Data: fi,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Result)
}
