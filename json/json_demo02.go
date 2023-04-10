package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	// 忽略空值字段“omitempty”
	Name string `json:"name,omitempty"`
	Age int64 `json:"age,omitempty"`
	// “-”忽略某字段
	Weight float64 `json:"-"`
}
func main() {
	p1 := Person{
		//Name: "Tom",
		Age: 8,
		//Weight: 30.0,
	}
	// json-> str
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("Marshal error: %v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

	// str->json
	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("Unmarshal error: %v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
