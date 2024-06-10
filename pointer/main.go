package main

import "fmt"

func main() {
	n := 10
	number := &n // & digunakan untuk mendapatkan alamat memori dari variable

	fmt.Println(number)
	fmt.Println(*number) // 10
	// * digunakan untuk mendapatkan nilai asli dari alamat memori

	// merubah variable "n" akan ikut mengubah variable pointer yang
	// mereferensi ke "n"
	n += 10

	fmt.Println(*number) // 20

	angka := 10
	ubahNilai(&angka)
	fmt.Println(angka)

	// var nama *string
	nama := new(string)
	fmt.Println(*nama) // nil

	// var fruits []string
	// if fruits == nil || len(fruits) <= 0 {
	// 	fmt.Println("fruits kosong")
	// }
}

func ubahNilai(n *int) { // argument by reference
	*n += 10
}
