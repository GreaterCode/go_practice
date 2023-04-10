/**
@author:admin
@date:2023/2/10
@note： json格式数据反序列化为map[string]interface{}
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main()  {
	type student struct {
		ID int `json:"id"`
		Name string `json:"name"`
	}

	s := student{ID: 123456789, Name: "Tom"}
	b, _ := json.Marshal(s)
	var m map[string]interface{}

	json.Unmarshal(b, &m)
	fmt.Println("id:%#v\n", m["id"])
	fmt.Println("id type:%T\n", m["id"])

	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	decoder.Decode(&m)
	fmt.Println("id:%#v\n", m["id"])
	fmt.Println("id type:%T\n", m["id"])
}
