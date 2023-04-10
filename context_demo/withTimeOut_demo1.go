package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go GenUsers(ctx, &wg)
	wg.Wait()
	fmt.Println("生成幸运用户")
}

func GenUsers(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("开始生成幸运用户")
	users := make([]int, 0)
	guser:for{
		select {
			case <-ctx.Done():
				fmt.Println(users)
				wg.Done()
				break guser
				return
		default:
			users = append(users, genUsersID(1000, 100000))
		}
	}
}

func genUsersID(min int, max int) int {
	return rand.Intn(max-min)+min
}


