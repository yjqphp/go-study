package main

import (
	"fmt"
	"math"
)

/**
	Go的if语句与for循环类似，表达式外无需小括号（），而大括号{}则是必须的
 */
func sqrt(x float64) string  {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func main()  {
	fmt.Println(sqrt(2),sqrt(-4))
}
