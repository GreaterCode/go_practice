package main

import (
	"encoding/json"
	"fmt"
)

type User1 struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User1  //匿名嵌套
	Password *struct{} `json:"password,omitempty"`
}
func main() {
	u1 := User1{
		Name:     "七米",
		Password: "123456",
	}
	b, err := json.Marshal(PublicUser{User1: &u1})
	if err != nil {
		fmt.Printf("json.Marshal u1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)  // str:{"name":"七米"}
}
