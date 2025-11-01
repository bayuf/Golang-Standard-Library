package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Silfi", "Bayu", "Firman", "Dian", "Putri"}
	number := []int{20, 1, 4, 2, 4, 6, 7, 3, 2, 8, 10}

	fmt.Println(slices.Index(names, "Silfi"))
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(number))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(number))
	fmt.Println(slices.Contains(names, "Silfi"))

}
