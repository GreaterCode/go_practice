package main

import (
	"encoding/json"
	"fmt"
 )

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}

func testFloat64() {
	var num1 float64 =2345.67
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误")
	}
	fmt.Printf("num1序列化后:%v\n", string(data))

	var num2 float64
	err = json.Unmarshal([]byte(data), &num2)
	if err != nil {
		fmt.Printf("反序列化error:%v\n", err)
	}
	fmt.Printf("反序列化后结果:%v\n", num2)
}

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "beijing"

	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"]= "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"chengdou", "xian"}

	slice = append(slice, m2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("marshal error: %v\n", err)
	}
	fmt.Printf(" slice 序列化后: %v\n",string(data))



}

func testMap() {
	var test_map map[string]interface{}
	test_map = make(map[string]interface{})
	test_map["name"] = "monkey"
	test_map["age"] = 20
	test_map["address"] = "hongyundong"

	data, err := json.Marshal(test_map)
	if err != nil {
		fmt.Printf("marshal error: %v\n",err)
	}

	fmt.Printf("map 序列化后:%v\n", string(data))

	var test_map2 map[string]string
	err = json.Unmarshal([]byte(data), &test_map2)
	if err != nil {
		fmt.Printf("map unmarshal error:%v\n", err)
	}
	fmt.Printf("map unmarshal result:%v\n", test_map2)
}

type Monster struct {
	Name string `json:"monster_name"` //反射机制
	Age int `json:"monster_age"`
	Birthday string //....
	Sal float64
	Skill string
}

func testStruct() {
	monster := Monster{
		Name: "panda",
		Age: 30,
		Birthday: "2022-1-22",
		Sal: 8000.0,
		Skill: "kunfu",
	}
	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Printf("struct marshal error: %v\n", err)
	}
	fmt.Printf("struct 序列化后: %v\n", string(data))

	var monster2 Monster
	err = json.Unmarshal([]byte(data),&monster2)
	if err != nil {
		fmt.Printf("struct unmarshal error: %v\n", err)
	}
	fmt.Printf("struct unmarshal result: %v\n", monster2.Name)

}
