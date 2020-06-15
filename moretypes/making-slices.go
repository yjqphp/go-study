package main

import "fmt"

/**
	用 make 创建切片

	切片可以用内建函数make来创建，
	这也是你创建动态数组的方式

	make函数会分配一个元素为零值的数组并返回一个引用了它的切片：
	a := make([]int, 5)

	要指定它的容量，需向make 传入第三个参数

	b := make([]int, 0, 5) //len(b)=5 cap(b) =5

	b = b[:cap(b)] //len(b)=5 cap(b)=5
	b = b[1:]      //len(b) cap(b) = 4
 */
func main()  {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)


}

func printSlice(s string, x []int)  {
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}