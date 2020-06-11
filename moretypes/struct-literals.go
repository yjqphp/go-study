package main

import "fmt"

/**
	结构体文法

	结构体文法通过直接列出字段的值来新分配一个结构体。

	使用Name: 语法可以列出部分字段。（字段名的顺序无关）

	特殊的前缀 & 返回一个指向结构体的指针
 */

type Vertex struct {
	X,Y int
}

var(
	v1 = Vertex{1,2} //创建一个Vertex类型的结构体
	v2 = Vertex{X:1}      //Y:0别隐式的赋予
	v3 = Vertex{}		  //X:0 Y:0
	p  = &Vertex{1,2} //创建一个 *Vertex类型的结构体（指针）
)

func main()  {
	fmt.Println(v1,p,v2,v3)
}
