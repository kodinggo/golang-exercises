package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Hello ")
	fmt.Print("World \n\n")

	fmt.Println("Hello")
	fmt.Println("World")

	log.Fatal("ini fatal")
	os.Exit(1)

	fmt.Printf("\n Nama saya %s, umur %d \n", "John", 10)

	var nama string
	fmt.Scan(&nama)
	fmt.Printf("Hello %s \n", nama)

	fmt.Println(fmt.Sprintf("I'm %d years old", 20))

	log.Println("test")

	err, _ := connectDB()
	if err != nil {
		log.Println(err)
	}
}

func connectDB() (error, bool) {
	// proses koneksi

	return errors.New("gagal koneksi"), false
}
