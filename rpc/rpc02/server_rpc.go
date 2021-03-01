package main

import (
	_ "fmt"
	"log"
	"net"
	"net/rpc"
)
//规范---1.服务名字定义
const HelloServiceName = "path/to/pkg.HelloService"

//规范---2.服务要实现的详细方法列表
type HelloServiceInterface = interface {
	Hello(request string,reply *string) error
}
//规范---3.注册改服务类型的函数
func RegisterHelloService(svc HelloServiceInterface) error {
	return  rpc.RegisterName(HelloServiceName,svc)

}

type HelloService struct {

}

func (p *HelloService ) Hello(request string,retry *string) error {
	*retry = "hello:" + request
	return  nil
}
func main()  {
	_ = RegisterHelloService(new(HelloService))
	listener,err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn,err := listener.Accept()
	if err!=nil {
		log.Fatal("Accept error:",err)
	}
	rpc.ServeConn(conn)

}
