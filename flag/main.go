package main

import (
	"flag"
	"fmt"
	"time"
)

// 定义命令行参数对应的变量，这四个变量都是指针类型
var cliName = flag.String("name", "John", "Input Your Name")
var cliAge = flag.Int("age", 22, "Input Your Age")
var cliGender = flag.String("gender", "male", "Input Your Gender")
var cliPeriod = flag.Duration("period", 1*time.Second, "sleep period")
 

// 定义一个值类型的命令行参数变量，再Init()函数中对其初始化，因此命令行参数对应变量的定义和初始化是可以的
var cliFlag int

func Init() {
	flag.IntVar(&cliFlag, "flagname", 1234, "Just for demo")
 }
func main() {
	// 初始化变量 cliFlag
	Init()
	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()
	//flag.Parsed()

	// flag.Args() 函数返回没有被解析的命令行参数
	// func NArg() 函数返回没有被解析的命令行参数的个数
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	// 输出命令行参数
	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("period=", *cliPeriod)
	fmt.Println("flagname=", cliFlag)

	fmt.Println("args num:", flag.NFlag())

	//$ go run main.go  --age=100
	//args=[], num=0
	//name= John
	//age= 100
	//gender= male
	//period= 1s
	//flagname= 1234
	//args num: 1

	//$ go run main.go  --help
	//Usage of C:\Users\admin\AppData\Local\Temp\go-build671458550\b001\exe\main.exe:
	//-age int
	//Input Your Age (default 22)
	//-flagname int
	//Just for demo (default 1234)
	//-gender string
	//Input Your Gender (default "male")
	//-name string
	//Input Your Name (default "John")
	//-period duration
	//sleep period (default 1s)

}