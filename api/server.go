package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	g := gin.Default()
	g.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "success"})
	})

	ser := http.NewServer(
		server.Name("test-service"),
		server.Address("127.0.0.1:8081"),
	)

	err := ser.Handle(ser.NewHandler(g))
	if err != nil {
		log.Println(err)
		return
	}

	service := micro.NewService(
		micro.Server(ser),
		micro.Registry(consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))),
	)

	service.Init()
	service.Run()
}
