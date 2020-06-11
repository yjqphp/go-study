package main

import (
	"fmt"
	"math"
)

/*
	表达式T(v) 将值v转换类型为T
	一些关于数值的转换：
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	或者,更加简单的形式

	i:=42
	f:=float64(i)
	u:=uint(f)

	与C不同的是,Go在不同类型的项之间赋值时需要显示转换
 */
func main(){
	var x,y int = 3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x,y,z)
}
