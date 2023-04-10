package main

import (
	"bytes"
	"fmt"
)

func main() {

	// 任意类型的可变参数
	var v1 int =1
	var v2 int64 =234
	var v3 string = "hello world"
	var v4 float64 = 1.2345
	MyPrintf(v1, v2, v3, v4)

	// 遍历可变参数列表——获取每一个参数的值
	fmt.Println(joinStrings("pig ", "and ", "rat "))
	fmt.Println(joinStrings("hammer", "mom", "hawk"))

	// 在多个可变参数函数中传递参数
	MyPrint(1, 2, 3)

	// 将不同类型变量通过printTypeValue()打印出来
	fmt.Println(printTypeValue(100, "str", true))
}

func printTypeValue(slist ...interface{}) string {
	// 字节缓冲作为快速字符串连接
	var b bytes.Buffer

	// 遍历参数
	for _, s := range slist {
		str := fmt.Sprintf("%v", s)
		var typeString string

		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}

		b.WriteString("value: ")
		b.WriteString(str)
		b.WriteString(" type: ")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}

func MyPrint(slist ... interface{} ) {
	rawPrint(slist...)
}

func rawPrint(rawList ...interface{}) {
	for _, a := range rawList {
		fmt.Println(a)
	}
}

func joinStrings(slist ...string) interface{} {
	var b bytes.Buffer
	for _, s:= range slist {
		b.WriteString(s)
	}
	return b.String()
}

func MyPrintf(args ... interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg,"is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}
