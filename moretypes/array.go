package main

import "fmt"

/**
	类型[n]T 表示拥有n个T类型的值的数组

	表达式
	var a [10]int

	会将变量a声明为拥有10个整数的数组

	数组的长度是其类型的一部分，因此数组不能修改变大小。这看起来是个限制,
	不过没关系,Go提供了更加便利的方式来使用数组
 */
func main()  {
	var a [2]string
	a[0] = "hello"
	a[1] = "World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
}
