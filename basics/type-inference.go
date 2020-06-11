package main

import "fmt"

/**
	类型推导
	在声明一个变量而不指定其类型时（即使使用不带类型的:=语法或var =表达式语法），变量的类型由右值推导得出

	当右值声了类型时，新变量的类型与其相同

	var i int
	j := i // j也是一个int

	不过当右边包含未指明类型的数值常量是int,float64或者complex128了,这取决于常量的精度：

	i :=42   //int   //int
	f := 3.142 //float  //float64
	g := 0.867 + 0.5i   //complex128
 */
func main(){
	v:= 0.867+0.5i
	fmt.Printf("v is of type %T\n", v)
}