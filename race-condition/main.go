package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1000)

	var x int // default 0

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			mu.Lock() // Memblock agar tidak ada goroutine lain yang mengakses variable "x"
			x++
			mu.Unlock() // Membuka "block", dan mengizinkan goroutine lain mengakses variable "x"
		}()
	}

	wg.Wait()

	fmt.Printf("x=%d\n", x)
}
