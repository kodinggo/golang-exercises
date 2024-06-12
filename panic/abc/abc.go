package abc

import (
	"fmt"
	"log"
)

func Blah() {
	defer func() {
		// Recover bertindak untuk mencegah terjadi error panic
		if err := recover(); err != nil {
			log.Println("Ada panic nih")
		}
	}()

	// Ini akan menyebabkan panic
	var abc *int
	fmt.Println(*abc)
}
