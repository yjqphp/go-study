package main

/**
	同for一样，if语句可以再条件表达式钱执行一个简单的语句。
	该语句声明的变量作用域仅在if之内
 */
import (
	"fmt"
	"math"
)

func pow(x,n,lim float64) float64  {
	if v := math.Pow(x,n); v< lim {
		return v
	}
	return lim
}

func main()  {
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
	)
}
