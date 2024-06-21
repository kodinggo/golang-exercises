package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3) // Inisialisasi total delta

	msgCh := make(chan string)

	go func() {
		defer wg.Done() // Mengurangi delta -1
		msgCh <- "Hello from GR 1"
	}()

	go func() {
		defer wg.Done() // Mengurangi delta -1
		msgCh <- "Hello from GR 2"
	}()

	go func() {
		defer wg.Done() // Mengurangi delta -1
		msgCh <- "Hello from GR 3"
	}()

	go func() {
		wg.Wait() // Menunggu sampai delta menjadi 0
		close(msgCh)
	}()

	for msg := range msgCh {
		fmt.Println(msg)
	}

	fmt.Println("Next program is executed")
}
