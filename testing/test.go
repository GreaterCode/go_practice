package main
import (
	"fmt"
	"os"``w0
)

func test_fmt_print()  {
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err :=  os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错")
		return
	}
	name := "123456"
	fmt.Fprintf(fileObj, "%s", name)
}

//func test_scan() {
//	var (
//		name    string
//		age     int
//		married bool
//	)
//	fmt.Scanln(&name, &age, &married)
//	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
//}

func main() {
	 a := [3]int{1, 2,3}
	 fmt.Println(arraySum(a))
	 test_fmt_print()
	 //test_scan()
}

func arraySum(a [3]int) interface{} {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}


