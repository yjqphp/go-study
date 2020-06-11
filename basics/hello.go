package main

import (
	"fmt"
)

var x,y int
var (
	ga int
	gb bool
)
var gc,gd int  = 1,2
var ge,gf = 123,"hello"

//这种不带声明格式的只能在函数体中出现

//gg,gh := 123,"hello"
func main()  {
	var i int
	var f float32
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var d = true
	fmt.Println(d)

	var intVal int

	// intVal :=1  这时会产生编译错误

	//此时不会产生编译错误，因为有声明新的变量，因为 :=是一个声明语句
	intVal,intVal1:=1,2
	fmt.Println(intVal)
	fmt.Println(intVal1)

	//多变量声明
	//1.类型相同多个变量，非全局变量

	//var vname1,vname2,vname3 int
	//vname1,vname2,vname3 = v1,v2,v3
	//
	//var vname1, vname2, vname3 = v1,v2,v3
	//
	//vname4,vname5,vname6 := v4,v5,v6

	//这种因式分解字的写法一般用于声明全局变量
	//var(
	//	vname1 int
	//	vname2 int
	//)

	gg,gh := 123,"hello"
	println(x,y,ga,gb,gc,gd,ge,gf,gg,gh)
}
