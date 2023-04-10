/**
@author:admin
@date:2023/3/31
@note time和ticker使用对比
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup
	timer1 := time.NewTimer(2 * time.Second)
	ticker1 := time.NewTicker(2 * time.Second)

	wg.Add(1)
	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<- t.C
			fmt.Println("exec ticker", time.Now().Format("2021-01-02 15:03:03"))
		}
	}(ticker1)

	wg.Add(1)
	go func(t *time.Timer) {
		defer wg.Done()
		for  {
			<-t.C
			fmt.Println("exec timer", time.Now().Format("2021-01-02 15:03:03"))
			t.Reset(2 * time.Second)
		}

	}(timer1)
	wg.Wait()
}