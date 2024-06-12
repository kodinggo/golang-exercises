package main

import (
	"fmt"
	"log"
)

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("ada error panic nih:", err)
	// 	}
	// }()

	// if !isEqual(1, 2) {
	// 	panic("dua nilai tidak sama")
	// }

	value := getValueSafely([]string{"satu", "dua"}, 2)
	fmt.Println(value)
}

func isEqual(n1, n2 int) bool {
	return n1 == n2
}

func getValueSafely(slice []string, index int) (value string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic error: ", err)
			value = "index tidak ditemukan"
		}
	}()

	return slice[index]
}
