package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 10
	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a)

	var b int = 077
	fmt.Printf("%d\n", b)

	var c int = 0xff
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)

	// 打印浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	fmt.Printf("%f\n", math.MaxFloat32)

	// 打印复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Printf("%v\n", c1)
	fmt.Printf("%v\n", c2)

	var c4 string = "qq"
	var b2 string = "cc"
	var c3 string = ""
	fmt.Printf("%d\n", len(c4))
	c3 = c4 + b2
	fmt.Printf("%s\n", c3)
}
