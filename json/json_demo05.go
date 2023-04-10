package main

/**
 * 优雅处理字符串格式的数字
 * 前端在传递来的json数据中可能会使用字符串类型的数字，这个时候可以在结构体tag中添加string来告诉json包从字符串中解析相应字段的数据：
 */

import (
	"encoding/json"
	"fmt"
)

type Card struct {
	ID int64 `json:"id,string"`
	Score float64 `json:"score,string"`
}

func main() {
	jsonStr1 := `{"id":"1234567","score":"88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("Error unmarsh jsonStr1 failed: %v\n", err)
		return
	}
	fmt.Printf("c1: %v\n", c1)
}
