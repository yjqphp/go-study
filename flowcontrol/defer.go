package main

import "fmt"

/**
	defer 语句会将函数推迟到外层函数返回之后执行。

	推迟调用的函数其参数会立即求值,但直到最外层函数返回该函数前
	该函数都不会被调用
 */
func main()  {
	defer fmt.Println("world")

	fmt.Println("hello")
}
