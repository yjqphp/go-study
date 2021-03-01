package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//1.通过rpc.Dial拨号RPC服务
	client,err := rpc.Dial("tcp","localhost:1234")
	if err != nil {
		log.Fatal("dialing",err)
	}
	var reply string
	//2.通过client.Call（）调用具体的RPC方法。
	//第一个参数是用点号链接的RPC服务名字和方法名字，
	//第二个和第三个参数分别是定义RPC方法的俩个参数
	err = client.Call("HelloService.Hello","hello",&reply)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
