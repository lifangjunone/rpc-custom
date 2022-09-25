package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc-custom/rpc_http_json/service"
)

var _ service.HelloService = (*HelloService)(nil)

type HelloService struct{}

type RPCReadwriteCloser struct {
	io.Writer
	io.ReadCloser
}

func NewRPCReadWriteCloser(w http.ResponseWriter, r *http.Request) *RPCReadwriteCloser {
	return &RPCReadwriteCloser{w, r.Body}
}

func (h *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello %s", request)
	return nil
}

func (h *HelloService) Add(request service.AddRequest, response *int) error {
	*response = request.A + request.B
	return nil
}

func main() {
	rpc.RegisterName(service.SERVICE_NAME, &HelloService{})
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		rpc.ServeCodec(jsonrpc.NewServerCodec(NewRPCReadWriteCloser(w, r)))
	})
	http.ListenAndServe(":1234", nil)
}
