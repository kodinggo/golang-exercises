package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	results := make([]int, len(numbers))

	for i, number := range numbers {
		go calculateSquare(number, i, results)
	}

	fmt.Println(results)

	// pending exit program sampai input data dari keyboad + enter
	var s string
	fmt.Scan(&s)
}

func calculateSquare(n, i int, results []int) {
	fmt.Printf("The square of %d is %d\n", n, n*n)
	results[i] = n * n
}
