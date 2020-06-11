package main

import "fmt"
/**
	Go只有一种循环结构: for循环
	基本的for循环有三部分组成，它们用分号隔开：
		初始化语句：在第一次迭代前执行
		条件表达式：在每次迭代前求值
		后置语句： 在每次迭代的结尾执行
	初始化语句通常为一句短变量声明，该变量声明仅在for语句的作用域中可见。
	一旦条件表达式的布尔值为false,循环迭代就会终止

	注意： 和C、Java、javaScript之类的语言不同，Go的for语句后面的三个部分外没有小括号，
	大括号{}则是必须的。


 */
func main()  {
	sum := 0
	for i:=0;i<10;i++ {
		sum += i
	}
	fmt.Println(sum)
}