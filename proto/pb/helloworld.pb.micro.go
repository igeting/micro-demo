// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: helloworld.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for HelloWorld service

func NewHelloWorldEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for HelloWorld service

type HelloWorldService interface {
	SayHello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type helloWorldService struct {
	c    client.Client
	name string
}

func NewHelloWorldService(name string, c client.Client) HelloWorldService {
	return &helloWorldService{
		c:    c,
		name: name,
	}
}

func (c *helloWorldService) SayHello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "HelloWorld.SayHello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloWorld service

type HelloWorldHandler interface {
	SayHello(context.Context, *Request, *Response) error
}

func RegisterHelloWorldHandler(s server.Server, hdlr HelloWorldHandler, opts ...server.HandlerOption) error {
	type helloWorld interface {
		SayHello(ctx context.Context, in *Request, out *Response) error
	}
	type HelloWorld struct {
		helloWorld
	}
	h := &helloWorldHandler{hdlr}
	return s.Handle(s.NewHandler(&HelloWorld{h}, opts...))
}

type helloWorldHandler struct {
	HelloWorldHandler
}

func (h *helloWorldHandler) SayHello(ctx context.Context, in *Request, out *Response) error {
	return h.HelloWorldHandler.SayHello(ctx, in, out)
}
