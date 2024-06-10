package main

import "fmt"

// Variabel "var" dapat dibuat di luar fungsi
var dbName = "test"

func main() {
	// variable
	var angka int = 10
	var nama string = "John"

	fmt.Println(angka)
	fmt.Println(nama)

	// variable tanpa tipe data
	var age = 20

	fmt.Println(age)
	fmt.Printf("%T\n", age) // untuk mendapatkan tipe data

	// variable tanpa value
	var address string
	var isSignup bool
	var number int

	fmt.Println(address)
	fmt.Println(isSignup)
	fmt.Println(number)

	// shorthand variabel (hanya bisa dibuat di dalam fungsi)
	firstName := "David"
	fmt.Println(firstName)

	var tinggi, lebar = 10, 20
	fmt.Println(tinggi, lebar)

	nama, umur := "Andi", 20
	fmt.Println(nama, umur)

	const dbName = "test"
	fmt.Println(dbName)
}
