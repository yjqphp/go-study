package main

//go语言常量
func main() {
	//const LENGTH int = 10
	//const WIDTH int = 5
	//var area int
	////多重赋值
	//const a,b,c = 1,false,"str"
	//
	//area = LENGTH * WIDTH
	//fmt.Printf("面积为：%d", area)
	//println()
	//println(a,b,c)

	//const (
	//	Unknow = 0
	//	Female = 1
	//	Male = 2
	//)

	//const(
	//	a = "abc"
	//	b = len(a)
	//	c = unsafe.Sizeof(a)
	//)
	//
	//println(a,b,c)

	const(
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值， iota += 1
		e        //"ha"  iota += 1
		f = 100  //iota += 1
		g        //100 iota += 1
		h = iota //7,恢复计数
		i        //8
	)
	println(a,b,c,d,e,f,g,h,i)
}
