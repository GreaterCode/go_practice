package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)
var logg * log.Logger
 func main() {
	logg := log.New(os.Stdout, "", log.Ltime)
	timeoutHandler()
	logg.Printf("ending")
}

func timeoutHandler() {
	ctx, cancel := context.WithDeadline(context.Background(),time.Now().Add(time.Second * 5))
	// ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	go doTimeOutStuff(ctx)
	time.Sleep(10 *time.Second)
	cancel()
}

func doStuff1(ctx context.Context) {
	var i = 1
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Printf("work %d seconds\n",i)
		}
		i++
	}
}

func doTimeOutStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		if deadline, ok := ctx.Deadline(); ok {
			fmt.Printf("deadline set:%d\n", deadline.Format("2006.01.02 15:04:05"))
 			if time.Now().After(deadline) {
				fmt.Printf(ctx.Err().Error())
				return
			}
		}
		select {
		case <-ctx.Done():
			fmt.Printf("done\n")
			return
		default:
			fmt.Printf("work\n")
		}
	}
}

