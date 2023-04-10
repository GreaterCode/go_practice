package main

import "fmt"

//把函数名称作为值保存到变量
func main() {
	var f func()
	f = fire
	f()
}

func fire()  {
	fmt.Println("fire")
}
