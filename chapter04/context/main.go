package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func f(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// 中断
			return
		default:
		}
		fmt.Println("goroutine: 処理")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(5*time.Second),
	)
	defer cancel()
	go f(ctx, &wg)

	// time.Sleep(10 * time.Second)
	// cancel()
	wg.Wait()
}
