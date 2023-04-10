package main

import (
	"fmt"
	"time"
)

const SecondsPerMinute = 60
const SecondsPerHour = 60 * SecondsPerMinute
const SecondsPerDay = 24 * SecondsPerHour

func main() {
	start := time.Now()

	//打印参数
	fmt.Println(resolveTime(1000))
	//只打印小时和分钟
	_, hour, minute := resolveTime(1800)
	fmt.Println(hour, minute)
	//只打印天
	day, _, _ := resolveTime(90000)
	fmt.Println(day)

	sum := 0
	for i :=0; i < 10000; i++ {
		sum++
	}

	//  time.Since(start) <=> time.Now().Sub(start)
	//elaped := time.Since(start)
	elaped := time.Now().Sub(start)
	fmt.Println("slaped: ", elaped)
 }

func resolveTime(seconds int) (day int, hour int, minute int)  {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds /SecondsPerMinute
	return
}