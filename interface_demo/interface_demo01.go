package main

import "fmt"

func main() {
	var any interface{}
	any = 1
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	any = 2
	fmt.Println(any)

	var b int = any.(int)
	fmt.Println(b)



}
