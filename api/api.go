package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
	"io"
	"log"
	"net/http"
)

func main() {
	service, err := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	).GetService("test-service")
	if err != nil {
		log.Fatalln(err)
	}

	next := selector.RoundRobin(service)
	node, err := next()
	if err != nil {
		log.Fatalln(err)
	}

	target := fmt.Sprintf("http://%s/test", node.Address)
	res, err := http.Post(target, "application/json", nil)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(data))
}
