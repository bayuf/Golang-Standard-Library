package main

// https://pkg.go.dev/fmt

import "fmt"

func main() {
	name := "Bayu Firmansyah"
	age := 25

	// Pakai format Println
	fmt.Println("Halo nama saya", name, "Umur saya", age, "Tahun")

	// Pakai format Printf
	fmt.Printf("Halo nama saya %s, umur saya %d Tahun\n", name, age)
}
