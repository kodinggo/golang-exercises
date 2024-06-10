package main

import "fmt"

type Label string

// Method dari type Label
func (l Label) convertToStr() string {
	return string(l)
}

func main() {
	var label Label = "abc"
	cetakLabel(label)

	fmt.Println(label.convertToStr()) // konversi Label ke string

	str := Label("test")
	cetakLabel(str)
}

func cetakLabel(l Label) {
	fmt.Println(l)
}
