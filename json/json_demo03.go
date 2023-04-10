/**
@author:admin
@date:2023/2/10
@n忽略嵌套结构体空值字段
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email,omitempty"`
	Hobby []string `json:"hobby,omitempty"`
	//Profile
	*Profile `json:"profile,omitempty"`
}

type Profile struct {
	Website string `json:"site"`
	Slogan  string `json:"slogan"`
}

func main() {
	u1 := User{
		Name:  "七米",
		Hobby: []string{"足球", "双色球"},
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}
