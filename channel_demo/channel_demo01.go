package main

import (
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]uint64, 10)
	lock sync.Mutex
)


func test(n int)  {
	var res uint64 = 1
	for i := 1;i < n;i++ {
		res *= uint64(i)
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20;i++ {
		go test(i)
	}
	lock.Lock()
	for i,v := range myMap {
		fmt.Printf("map[%d]=%d\n",i, v)
	}
	lock.Unlock()
}
