package main

import (
	"fmt"
	"strings"
)

// 操作与数据分离的设计技巧
func main() {
	// 待处理字符串
	list := []string {
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	// 处理链
	chain := []func(string) string {
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	//处理字符串
	 StringProcess(list, chain)

	// 输出处理好的字符串
	for _, str := range list {
		fmt.Println(str)
	}
}

func StringProcess(list []string, chain []func(string) string) {
	for index, str := range list {
		result := str
		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
	}
}

func removePrefix(s string) string {
	return strings.TrimPrefix(s, "go")
}
