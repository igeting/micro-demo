package main

import (
	"context"
	"github.com/asim/go-micro/plugins/client/http/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"log"
)

func main() {
	cli := http.NewClient(
		client.Selector(
			selector.NewSelector(
				selector.SetStrategy(selector.Random),
			),
		),
		client.Registry(
			consul.NewRegistry(
				registry.Addrs("127.0.0.1:8500"),
			),
		),
		client.ContentType("application/json"),
	)

	req := cli.NewRequest("test-service", "/test", nil)

	var res map[string]interface{}
	err := cli.Call(context.TODO(), req, &res)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res)
}
