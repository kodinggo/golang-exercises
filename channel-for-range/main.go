package main

import (
	"fmt"
	"time"
)

func main() {
	msgCh := make(chan string)

	go func() {
		for i := 1; i <= 5; i++ {
			msgCh <- fmt.Sprintf("ini pesan ke %d", i)
		}
		close(msgCh) // memberi signal ke proses for-range bahwa tidak ada lagi groutine yg mengiirmkan data ke channel
	}()

	for msg := range msgCh {
		fmt.Println(msg)
	}

	time.Sleep(1 * time.Second)
}
