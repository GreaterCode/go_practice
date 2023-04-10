package main

import "fmt"

// 参数传递效果测试

type InnerData struct {
	a int
}

type Data struct {
	complax []int
	instance InnerData
	ptr *InnerData
}

func main() {
	in := Data{
		complax: []int{1,2,3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}

	// 输入结构体成员
	fmt.Printf("in value: %+v\n",in)
	fmt.Printf("in  ptr: %p\n", &in)

	// 传入结构体，返回同类型的结构体
	out := paasByValue(in)
	fmt.Printf("out value: %+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)

}

func paasByValue(inFunc Data)  Data {
	fmt.Printf("inFunc value:%+v\n", inFunc)
	fmt.Printf("inFunc ptr:%p\n", &inFunc)

	return inFunc
}
