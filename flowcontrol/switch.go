package main

import (
	"fmt"
	"runtime"
)

/**
	switch
	Go只运行选定的case，而非之后所有的case。实际上，Go自动提供在这些
	语言中每个case后面所需的break语句。除非以fallthrough语句结束，否则
	分支会自动终止。Go的另一个点重要的不同于switch的case无需为常量，且取值不必为整数。
 */
func main(){
	fmt.Print("Go runs on ")
	switch os:= runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n",os)
	}
}
