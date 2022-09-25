package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc-custom/rpc_tcp_json/service"
)

var _ service.HelloService = (*HelloService)(nil)

type HelloService struct{}

func (h *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello %s", request)
	return nil
}

func main() {
	rpc.RegisterName(service.SERVICE_NAME, &HelloService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err.Error())
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
