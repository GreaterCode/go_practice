/**
@author:admin
@date:2023/2/17
@note: select 实现优先级
*/

package main

import (
	"fmt"
	"time"
)

func worker(ch1, ch2 chan string, stopCh chan struct{}){

	for {
		select {
		case <- stopCh:
			return
		case test1 := <- ch1:
			fmt.Println(test1)
		case test2 := <- ch2:
			time.Sleep(time.Second * 2)
		priority:
			select {
			case test1 := <-ch1:
				fmt.Println(test1)
			default:
				break priority
			}
			fmt.Println(test2)
		}
	}
}

func main() {
	var ch1 = make(chan string)
	var ch2 = make(chan string)
	stopchan := make(chan struct{})
	go worker(ch1, ch2, stopchan )
	ch2 <- "ch2"
	time.Sleep(time.Second * 1)
	ch1 <- "1"

	time.Sleep(time.Second * 5)


}
