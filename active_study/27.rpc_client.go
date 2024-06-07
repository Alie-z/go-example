package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// HelloServiceName 服务的名字
const HelloServiceName = "path/to/pkg.HelloService"

// HelloServiceInterface 服务的接口
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

// RegisterHelloService 注册方法
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

// HelloServiceClient HelloService客户端，实现了HelloServiceInterface接口
type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

// DialHelloService 拨号HelloService服务
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

// Hello 实现了HelloService的接口方法
func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>> reply", reply)
}
