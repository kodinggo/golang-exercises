package main

import "fmt"

func main() {
	sayHello("Hello ", "Joko") // Memanggil fungsi

	// Memanggil fungsi dengan arguemnt variadic
	result := sum(1, 2, 3)
	fmt.Println(result)

	sum, dif := getSumAndDifferent(10, 20)
	fmt.Println("sum ", sum)
	fmt.Println("diferent ", dif)

	greet := func() {
		fmt.Println("Hello")
	}
	greet()

	result = multiple(1, 2, func(result int) int {
		return result
	})
	fmt.Println(result)

	s := say("John")
	fmt.Println(s())
}

// Fungsi dengan argument
func sayHello(greeting, nama string) {
	fmt.Println(greeting + nama)
}

// Fungsi dengan argument variadic
func sum(numbers ...int) int {
	var total int

	for _, n := range numbers {
		total += n
	}

	return total
}

// Fungsi dengan multiple argument
func getSumAndDifferent(n1, n2 int) (sum, dif int) {
	sum = n1 + n2

	if n1 > n2 {
		dif = n1 - n2
	} else {
		dif = n2 - n1
	}

	return
}

func multiple(n1, n2 int, callback func(int) int) int {
	result := n1 * n2
	return callback(result)
}

func say(s string) func() string {
	return func() string {
		return "Hello" + s
	}
}
