/**
@author:admin
@date:2023/2/13
@note:自定义解析时间字符串
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Comment struct {
	Content string
}

type Image struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type SendMessage struct {
	User string `json:"user"`
	Msg string `json:"msg"`
}


func main(){
	c1 := Comment{
		Content: "永远不要高估自己",
	}
	i1 := Image{
		Title: "赞赏码",
		URL:   "https://www.liwenzhou.com/images/zanshang_qr.jpg",
	}
	// struct -> json string
	b, err := json.Marshal(struct {
		*Comment
		*Image
	}{&c1, &i1})
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	// json string -> struct
	jsonStr := `{"Content":"永远不要高估自己","title":"赞赏码","url":"https://www.liwenzhou.com/images/zanshang_qr.jpg"}`
	var (
		c2 Comment
		i2 Image
	)
	if err := json.Unmarshal([]byte(jsonStr), &struct {
		*Comment
		*Image
	}{&c2, &i2}); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("c2:%#v i2:%#v\n", c2, i2)

	jsonStr = `{"sendMsg":{"user":"q1mi","msg":"永远不要高估自己"},"say":"Hello"}`
	var data map[string]json.RawMessage
	if err = json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Printf("json.Unmarshal error: %v\n", err)
		return
	}
	fmt.Printf("data: %v\n", data)

	var msg SendMessage
	if err := json.Unmarshal(data["sendMsg"], &msg); err != nil {
		fmt.Printf("json.Unmarshal error: %v\n",err)
		return
	}
	fmt.Printf("msg:%#v\n", msg)



}


