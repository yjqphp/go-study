package main

import "fmt"

/**
	指针

	Go拥有指针。指针保存了值的内存地址

	类型 *T 是指向T类型的指针。其零值为nil

	var p *int

	& 操作符会生成一个指向操作数的指针。
	i := 42
	p = &i

	* 操作符表示指针指向的底层值。

	fmt.Println(*p) //通过p
	*p = 21  // 通过指针p设置i

	这也就是通常所说的 "间接引用" 或 ”重定向“
	与C不同,Go没有指针运算
 */
func main()  {
	i,j := 42,2701

	p := &i  //指向i
	fmt.Println(*p) //通过指针读取i的值

	*p = 21 //通过指针设置i的值
	fmt.Println(i) //查看i的值

	p = &j
	*p= *p/37 //通过指针对j进行除法运算
	fmt.Println(j) //查看j的值
}
