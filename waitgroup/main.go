package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3) // Inisialisasi total delta

	go func() {
		fmt.Println("Hello from GR 1")
		wg.Done() // Mengurangi delta -1
	}()

	go func() {
		fmt.Println("Hello from GR 2")
		wg.Done() // Mengurangi delta -1
	}()

	go func() {
		fmt.Println("Hello from GR 3")
		wg.Done() // Mengurangi delta -1
	}()

	wg.Wait() // Menunggu sampai delta menjadi 0

	/*
	{
		detail artikel
		komentar
		author
	}
	*/

	fmt.Println("Next program is executed")
}
