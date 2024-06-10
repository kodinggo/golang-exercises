package main

import "fmt"

func main() {
	// array
	var fruits = [3]string{
		"Mangga",
		"Pisang",
		"jambu",
	}
	fmt.Println(fruits, len(fruits))

	// slice
	numbers := []int{
		1, 2, 3, 4, 5,
	}
	numbers = append(numbers, 6)
	fmt.Println(numbers, len(numbers))

	fmt.Println(numbers[0]) // 1

	blah := []any{1, "test", true, []int{1, 2, 3}}
	fmt.Println(blah)
}
