package service

const (
	SERVICE_NAME = "HelloService"
)

type AddRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type HelloService interface {
	Hello(request string, response *string) error
	Add(request AddRequest, response *int) error
}
