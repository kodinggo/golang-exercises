package main

import (
	"fmt"
	"time"
)

func main() {
	// menjalankan fungsi dengan goroutine
	// ini akan membuat proses eksekusi terjadi secara paralel
	go cetak("hello")
	go cetak("hi")
	go cetak("what's up")

	cetak("apa kabar?")

	time.Sleep(1 * time.Second)
}

func cetak(s string) {
	fmt.Println(s)
}
