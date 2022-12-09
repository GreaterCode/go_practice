package main

//iota是go语言的常量计数器，只能在常量的表达式中使用。
//
//iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

func main() {

	//const (
	//	n1 = iota //0
	//	n2        //1
	//	n3        //2
	//	n4        //3
	//)
	//fmt.Println(n1)
	//fmt.Println(n2)
	//fmt.Println(n3)
	//fmt.Println(n4)
	//0
	//1
	//2
	//3

	// 使用_跳过某些值
	//const (
	//	n1 = iota
	//	n2
	//	_
	//	n4
	//)
	//fmt.Println(n1)
	//fmt.Println(n2)
	//fmt.Println(n4)
	//
	//0
	//1
	//3

	//iota声明中间插队
	//const (
	//	n1 = iota //0
	//	n2 = 100  //100
	//	n3 = iota //2
	//	n4        //3
	//)
	//const n5 = iota //0
	//
	//fmt.Println(n1)
	//fmt.Println(n2)
	//fmt.Println(n3)
	//fmt.Println(n4)
	//
	//0
	//100
	//2
	//3

	// 定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<20表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
	//const (
	//	_ = iota
	//	KB = 1 << (10 * iota)
	//	MB = 1 << (10 * iota)
	//	GB = 1 << (10 * iota)
	//	TB = 1 << (10 * iota)
	//	PB = 1 << (10 * iota)
	//)
	//fmt.Println(KB)
	//fmt.Println(MB)
	//fmt.Println(GB)
	//fmt.Println(TB)
	//fmt.Println(PB)
	//
	//1024
	//1048576
	//1073741824
	//1099511627776
	//1125899906842624

	// 多个iota定义在一行
	//const (
	//	a , b = iota + 1 , iota + 2
	//	c, d
	//	e, f
	//)
	//
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
	//fmt.Println(e)
	//fmt.Println(f)
	//
	//1
	//2
	//2
	//3
	//3
	//4

}
