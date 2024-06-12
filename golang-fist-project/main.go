package main

import (
	"fmt" // Package bawaan

	"golang-first-project/greeting" // Package internal

	"github.com/google/uuid" // Package external
)

func main() {
	uuid, _ := uuid.NewV6()
	fmt.Println(uuid)

	greeting.Hello()
	greeting.Hi()
}
