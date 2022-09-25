package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc-custom/rpc_tcp_json/service"
)

var _ service.HelloService = (*HelloServiceClient)(nil)

type HelloServiceClient struct {
	client *rpc.Client
}

func NewHelloService(network, addr string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{
		client: client,
	}, nil
}

func (h *HelloServiceClient) Hello(request string, response *string) error {
	err := h.client.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, response)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	helloSvc, err := NewHelloService("tcp", ":1234")
	if err != nil {
		panic(err.Error())
	}
	var rsp string
	err = helloSvc.Hello("ldd", &rsp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rsp)
}
