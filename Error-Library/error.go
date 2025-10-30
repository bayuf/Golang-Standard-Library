package main

import (
	"errors"
	"fmt"
)

// https://pkg.go.dev/errors

var (
	ValidationErrors = errors.New("Validation Error")
	NotFoundErrors   = errors.New("Data Not Found")
)

func getById(id string) error {
	if id == "" {
		return ValidationErrors
	}

	if id != "Bayu" {
		return NotFoundErrors
	}

	return nil
}

func main() {
	data := getById("")

	if data != nil {
		if errors.Is(data, ValidationErrors) {
			fmt.Println("Validation Error :", data)
		} else if errors.Is(data, NotFoundErrors) {
			fmt.Println("Not Found Error :", data)
		} else {
			fmt.Println("Unknown Error")
		}
	}
}
