package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nama any = "Bambang"
	// tipe assertion
	strNama, ok := nama.(string)
	if ok {
		cetakStr(strNama)
	}

	// tipe conversion
	angka := "10"
	angkaNumber, err := strconv.Atoi(angka)
	if err != nil {
		panic("tipe data tidak didukung")
	}
	fmt.Println(angkaNumber)
}

func cetakStr(s string) {
	fmt.Println(s)
}
