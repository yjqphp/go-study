package main

import (
	"fmt"
	"time"
)

/**
	switch的求值顺序
	switch 的 case语句从上导下顺次执行，知道匹配成功时停止

	例如：
	switch i{
		case 0:
		case f():
	}
	在i==0时f不会被调用
 */
func main()  {
	fmt.Println("When's Saturday?")
	today :=time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")

	}
}
