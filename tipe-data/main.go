package main

import "fmt"

func main() {
	var x int = 10
	var y int64 = 10
	fmt.Println(x * int(y))

	char := 'A'
	fmt.Printf("unicode number=%d, tipe data=%T\n", char, char)

	str := "hello world"
	for _, char := range str {
		fmt.Printf("karakter=%c, unicode number=%d \n", char, char)
	}

	/*
		%d = konversi angka ke string
		%T = mendapat tipe data dan konversi ke string
		%c = konversi karakter ke string
	*/
}
