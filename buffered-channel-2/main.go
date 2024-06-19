package main

import "fmt"

func main() {
	totalBuffer := 5
	msgCh := make(chan int, totalBuffer)

	go func() {
		for i := 1; i <= totalBuffer; i++ {
			msgCh <- i
			fmt.Printf("%d is sent\n", i)
		}
		close(msgCh)
	}()

	for msg := range msgCh {
		fmt.Printf("%d is received\n", msg)
	}

	var i string
	fmt.Scan(&i)
}
