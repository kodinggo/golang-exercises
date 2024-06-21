package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1000)

	var x int64 // default 0

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&x, 1)
		}()
	}

	wg.Wait()

	fmt.Printf("x=%d\n", x)
}
