package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "abd"
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)

	s1 := `a
b\r\n\x00
c`
	println(s1)

	s2 := "hello" +
		"world"
	println(s2)

	s3 := "abcd"
	bs := []byte(s3)
	bs[1] = 'B'
	println(string(bs))

	u := "电脑"
	us := []rune(u)
	us[1] = '话'
	println(string(us))

	// 修改字符串
	s4 := "abc汉字"
	for i := 0; i < len(s4); i++ { //byte
		fmt.Printf("%c", s4[i])
	}
	println()
	for _, r := range s4 { // rune
		fmt.Printf("%c", r)
	}
	println()

	// 字符串替换
	s5 := "hello world world, world"
	res0 := strings.Replace(s5, "world", "golang", 2)
	res1 := strings.Replace(s5, "world", "golang", 1)
	res2 := strings.Replace(s5, "world", "golang", -1)

	// strings.Replace("原字符串", "被替换的内容", "替换的内容", 替换次数)
	fmt.Printf("res0 is %v\n", res0)
	fmt.Printf("res1 is %v\n", res1)
	fmt.Printf("res2 is %v\n", res2)

	arr1 := strings.Split("a,b,c", ",")
	fmt.Printf("%v\n", arr1) // [a b c]

	arr2 := strings.Split("amnmn", "m")
	fmt.Printf("%v\n", arr2) // [a n n]

	// 按照指定字符串切割原字符串，并切指定切割为几份，如果最后一个参数为0，那么范围是一个空数组
	arr3 := strings.SplitN("a,b,c", ",", 2)
	fmt.Println(arr3) // [a b,c]

	arr4 := strings.SplitN("a,nc,b,f", ",", 0)
	fmt.Printf("%v\n", arr4) // []

	arr5 := strings.SplitAfter("a,nc, b,f", "")
	fmt.Println(arr5)

	arr6 := strings.SplitAfterN("a,nc,b,f", ",", 2)
	fmt.Println(arr6)

	// 按照空格切割字符串, 多个空格会合并为一个空格处理
	arr7 := strings.Fields("a b c d    e")
	fmt.Println(arr7)

	// 将字符串切片按照指定连接符号转换为字符串
	sce := []string{"aa", "nn", "cc"}
	str1 := strings.Join(sce, "-")
	fmt.Println(str1)

	// 返回count个s串联的指定字符串
	str2 := strings.Repeat("123", 3)
	fmt.Println(str2)

	str3 := "123413445"
	str4 := strings.Replace(str3, "34", "bc", -1)
	fmt.Println(str4)

	//
	str5 := " 123 3454,,!!"
	str6 := strings.Trim(str5, " ")
	fmt.Println(str6)

	str7 := " 123 3454,,!!"
	str8 := strings.Trim(str7, " ")
	fmt.Println(str8)

	fmt.Println(utf8.RuneCountInString("人民"))
	fmt.Println(utf8.RuneCountInString("人民，good"))

	fmt.Println(strings.ContainsAny("in failure", "s g"))

}
