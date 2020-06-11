package main

import "fmt"

/**
	结构体字段可以通过结构体指针来访问

	如果我们有一个指向结构体的指针p,那么可以通过（*p）.X来访问其字段X.
	不过这么写太啰嗦了,所以语言也允许我们使用隐式间接引用，直接写p.X就可以
 */
type Vertex struct {
	X int
	Y int
}

func main()  {
	v := Vertex{1,2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}


