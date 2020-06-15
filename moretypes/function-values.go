package main

import (
	"fmt"
	"math"
)

/**
	函数值

	函数也是值。它们可以像其他值一样传递。

	函数值可以用作函数的参数或返回值
 */
func compute(fn func(float64,float64) float64) float64{
	return fn(3,4)
}
func main()  {
	hypot := func(x,y float64)  float64{
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12)) //math.Sqrt(5*5+12*12) = 13

	fmt.Println(compute(hypot))   //math.Sqrt(3*3 + 4*4) =  5

	fmt.Println(compute(math.Pow))   //math.Pow(3,4)=3^4  = 81
}
