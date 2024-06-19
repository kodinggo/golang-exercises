package main

import "fmt"

func main() {
	ch := make(chan int)
	msgCh := make(chan string)

	go func() {
		ch <- 1
		sender(msgCh, "hello")
	}()

	go func() {
		ch <- 2
		sender(msgCh, "hi")
	}()

	go func() {
		ch <- 3
		sender(msgCh, "good bye")
	}()

	var data int

	data = <-ch
	fmt.Println("data dari goroutine ke 1 = ", data)

	data = <-ch
	fmt.Println("data dari goroutine ke 2 = ", data)

	data = <-ch
	fmt.Println("data dari goroutine ke 3 = ", data)

	// data = <-ch
	// fmt.Println("data dari goroutine ke 3 = ", data)

	fmt.Println(receiver(msgCh))
	fmt.Println(receiver(msgCh))
	fmt.Println(receiver(msgCh))

	var s string
	fmt.Scan(&s)
}

func sender(msgCh chan<- string, msg string) {
	msgCh <- msg
}

func receiver(msgCh <-chan string) (msg string) {
	msg = <-msgCh
	return msg
}
