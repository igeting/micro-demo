package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	"log"
	"micro-demo/proto/pb"
)

func main() {
	service := micro.NewService()

	cli := pb.NewHelloWorldService("test-service", service.Client())

	res, err := cli.SayHello(context.Background(), &pb.Request{
		Name: "lulu",
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}
