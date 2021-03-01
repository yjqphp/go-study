package main

import (
	_ "fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {

}

/**
Hello方法必须满足go语言的rpc规则：
方法只能有俩个可序列化的参数，
其中第二个参数是指针类型，并且返回一个error类型
必须是公开的方法
 */
func (p *HelloService) Hello(request string,reply *string) error  {
	*reply = "Hello:" + request
	return nil
}
func main()  {
	//将对象类型中所有满足RPC规则的对象方法注册为RPC函数，
	//所有注册的方法会放在HelloService服务空间下。
	_ = rpc.RegisterName("HelloService", new(HelloService))
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
