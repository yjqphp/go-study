package main

import "fmt"

/**
	接口值

	接口也是值。它们可以像其它值一样传递。
	接口值可以用作函数的参数或返回值。
	在内部，接口值可以看做
 */
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M()  {
	fmt.Println(t.S)
}

type F float64

func (f F) M()  {
	fmt.Println(f)
}

func main()  {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v,%T)\n",i,i)
}