package main

import (
	"fmt"
	"math"
)

/**
	if和else

	在if的简短语句中声明的变量同样可以在任何对应的else块中使用
	（在main的fmt.Println调用开始前，两次对pow的调用均已执行并返回其各自的结果）
 */
func pow(x,n,lim float64) float64 {
	if v:=math.Pow(x,n); v<lim {
		return v
	} else {
		fmt.Printf("%g >=%g\n",v,lim)
	}
	//这里开始就不能使用vl
	return lim
}
func main()  {
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}
