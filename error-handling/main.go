package main

import (
	"error-handling/myerror"
	"fmt"
	"log"
)

func main() {
	slice := []string{"satu", "dua", "tiga"}
	val, err := getValue(slice, 3)
	// Cek apakah terjadi error
	// if err != nil {
	// 	log.Fatalf("terjadi error: %s", err.Error())
	// }
	if err != nil {
		switch err.(type) {
		case myerror.ErrorInvalidArgument:
			log.Fatalf("HTTP Code 400: %s", err.Error())
		case myerror.ErrorNotFound:
			log.Fatalf("HTTP Code 404: %s", err.Error())
		default:
			log.Fatalf("HTTP Code 500: %s", err.Error())
		}
	}

	fmt.Println(val)

}

func getValue(s []string, k int) (string, error) {
	if len(s)-1 < k {
		return "", myerror.NewErrorInvalidArgument("index is offset")
	}
	return s[k], nil
}
