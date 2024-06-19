package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Ini dari goroutine 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Ini dari goroutine 2"
	}()

	for i := 0; i < 2; i++ {
		// proses penerimaan data dari channel tidak saling tunggu
		select {
		case data := <-ch1:
			fmt.Println(data) // 2 detik
		case data := <-ch2:
			fmt.Println(data) // 1 detik
		}
	}

	var s string
	fmt.Scan(&s)
}
