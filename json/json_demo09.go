/**
@author:admin
@date:2023/2/13
@note: 自定义解析时间字符串
*/

package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"
var nilTime = (time.Time{}).UnixNano()

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

type Post struct {
	CreateTime CustomTime `json:"create_time"`
}

func (ct *CustomTime) MarshalJSON() ([]byte, error)  {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}
func (ct *CustomTime) IsSet() bool  {
	return ct.UnixNano() != nilTime
}
func main()  {
	p1 := Post{CreateTime: CustomTime{time.Now()}}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal p1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	jsonStr := `{"create_time":"2020-04-05 12:25:42"}`
	var p2 Post
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
