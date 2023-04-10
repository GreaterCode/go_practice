/**
@author:admin
@date:2023/2/13
@note:  json格式数据反序列化为map[string]interface{}
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	var m =  make(map[string]interface{}, 1)
	m["count"] = 1
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed: %v\n", err)
	}
	fmt.Printf("str:%#v\n", string(b))

	var m2 map[string]interface{}
	err = json.Unmarshal(b, &m2)
	if err != nil {
		fmt.Printf("unmarshal failed: %v\n", err)
		return
	}
	fmt.Printf("value:%#v\n", m2["count"])
	fmt.Printf("type:%T\n", m2["count"])

}