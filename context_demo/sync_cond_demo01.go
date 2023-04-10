package main

import (
	"log"
	"sync"
	"time"
)

var done = false

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	go write("writer1", cond)
	time.Sleep(time.Second * 6)
}

func write(s string, cond *sync.Cond) {
	println(s, "start writing")
	time.Sleep(time.Second)
	cond.L.Lock()
	done = true
	cond.L.Unlock()
	log.Println(s, "wakes all")
	/*
	 * Broadcast 唤醒所有等待条件变量 c 的 goroutine，无需锁保护
	 * Signal 只唤醒任意1个等待条件变量 c 的 goroutine，无需锁保护。 有点类似 Java 中的 Notify

	 */
	//cond.Broadcast()
	cond.Signal()
	time.Sleep(time.Second)
	cond.Signal()
	time.Sleep(time.Second)
	cond.Signal()
}

func read(s string, cond *sync.Cond) {
	cond.L.Lock()
	for !done {
		cond.Wait()
	}
	log.Println(s, "start reading")
	cond.L.Unlock()

}
