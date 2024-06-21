package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1000)

	var x int // default 0

	ch := make(chan int, 1) // buffer 1
	ch <- 5

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			<-ch // Memblock agar tidak ada goroutine lain yang mengakses variable "x"

			x++

			ch <- 5 // Membuka "block", dan mengizinkan goroutine lain mengakses variable "x"
		}()
	}

	wg.Wait()

	fmt.Printf("x=%d\n", x)
}
