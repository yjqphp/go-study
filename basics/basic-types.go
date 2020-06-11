package main

import (
	"fmt"
	"math/cmplx"
)

/**
	bool

	string

	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 unitptr

	byte //uint8别名
		表示一个Unicode码点

	float32 float64

	complex64 complex128

	变量声明也可以“分组”成一个语法块

	int uint和uintptr在32位系统上通常为32位宽,在64位系统上则为64位宽。
	当你需要一个整数值时应使用int类型，除非你有特殊的理由使用固定大小或者无符号的整数类型

 */

var (
	ToBe    bool = false
	MaxInt  uint64 = 1<<64-1
	z complex128 = cmplx.Sqrt(-5+12i)
)

func main()  {
	fmt.Printf("Type:%T Value:%v\n", ToBe, ToBe)
	fmt.Printf("Type:%T Value:%v\n", MaxInt, MaxInt)
	fmt.Printf("Type:%T Value:%v\n", z, z)
	
}
