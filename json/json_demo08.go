/**
@author:admin
@date:2023/2/13
@note:  json格式数据反序列化为map[string]interface{}
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	// map[string]interface{} -> json string
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 // int
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%#v\n", string(b))
	// json string -> map[string]interface{}
	var m2 map[string]interface{}
	// 使用decoder方式反序列化，指定使用number类型
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	err = decoder.Decode(&m2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("value:%v\n", m2["count"]) // 1
	fmt.Printf("type:%T\n", m2["count"])  // json.Number
	// 将m2["count"]转为json.Number之后调用Int64()方法获得int64类型的值
	count, err := m2["count"].(json.Number).Int64()
	if err != nil {
		fmt.Printf("parse to int64 failed, err:%v\n", err)
		return
	}
	fmt.Printf("type:%T\n", int(count)) // int

}