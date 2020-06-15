package main

import "fmt"

/**
	向切片追加元素
	https://go-zh.org/pkg/builtin/#append

	func append(s []T, vs ...T) []T

	append 的第一个参数s是一个元素类型为T的切片，其余类型为T的值将会
	追加到该切片的末尾。

	append 的结果是一个包含原切片所有元素加上新添加元素的切片。

	当s的底层数组太小,不足以容纳所有给定的值时,它就会分配一个更大的数组。返回的切片会指向这个新分配
	的数组。

	更多切片用法和本质： https://blog.go-zh.org/go-slices-usage-and-internals
 */

func main(){
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s,0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s,1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2,3,4)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}