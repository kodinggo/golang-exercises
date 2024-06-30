package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()

	go doSomething(ctx)

	// cancel ctx setelah 5 detik
	// time.AfterFunc(5*time.Second, func() {
	// 	fmt.Println("process has been 5 seocnds, timeout!!!")
	// 	cancel()
	// })

	var s string
	fmt.Scan(&s)
}

func doSomething(ctx context.Context) {
	for {
		if isCtxCanceled(ctx) {
			fmt.Println("stop process!")
			return
		}
		fmt.Println("dome some work...")
		time.Sleep(1 * time.Second)
	}
}

func isCtxCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return true
	default:
		return false
	}
}
