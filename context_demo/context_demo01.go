package main

import (
	"context"
	"log"
	"os"
	"time"
)

// var logg * log.Logger
func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("down")
}

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	// 10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}


// 每1秒就work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(1 * time.Second)
		select {
			case <-ctx.Done():
				logg.Printf("done")
				return
			default:
				logg.Printf("working %d s.\n", i)
		}
		i++
	}
}
