package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"log"
	"micro-demo/proto/pb"
)

func main() {
	service := micro.NewService()
	service.Init()

	cli := pb.NewHelloWorldService("test-service", service.Client())
	res, err := cli.SayHello(context.Background(), &pb.Request{Name: "lulu"})
	if err != nil {
		log.Println(err)
	}

	log.Println(res.Result)
}
