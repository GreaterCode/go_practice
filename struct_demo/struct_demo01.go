package main

import (
	"encoding/json"
	"fmt"
)

// 使用匿名结构体解析JSON数据

func main() {
	jsonData := getJSONData()
	fmt.Println(string(jsonData))

	// 只需要屏幕和指纹识别信息的结构和实例
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	// 反序列化得到数据
	json.Unmarshal(jsonData, &screenAndTouch)
	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", screenAndTouch)

	// 只需要电池和指纹识别信息的结构和实例
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}
	// 反序列化到batteryAndTouch
	json.Unmarshal(jsonData, &batteryAndTouch)
	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", batteryAndTouch)

}

type Screen struct {
	Size float32
	ResX, ResY int
}

type Battery struct {
	Capacity int
}

func getJSONData() []byte {
	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery {
			2910,
		},
		HasTouchID: true,
	}

	jsonData, _ := json.Marshal(raw)
	return jsonData
}
