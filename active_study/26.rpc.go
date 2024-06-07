package main

import (
	"log"
	"net"
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

// HelloService 服务对象类型，该服务可以对客户端提供Hello方法远程方法调用
type HelloService struct{}

// Hello 该方法用于实现打印功能
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
