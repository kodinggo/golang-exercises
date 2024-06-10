package main

import "fmt"

func main() {
	cetak("test")

	defer cetak("hello 1")
	defer cetak("hello 2")
	defer cetak("hello 3")

	cetak("world")
}

func cetak(s string) {
	fmt.Println(s)
}
