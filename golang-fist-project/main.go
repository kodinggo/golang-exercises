package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid, _ := uuid.NewV6()
	fmt.Println(uuid)
}
