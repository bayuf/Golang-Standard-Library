package main

import (
	"fmt"
	"strings"
)

func main() {

	name := " Bayu Firmansyah"

	// Contains untuk mencari karakter tertentu di dalam sebuah string
	fmt.Println(strings.Contains("bayu", "s"))
	fmt.Println(name)

	// Trim Menghilangkan space di kiri dan kanan string
	fmt.Println(strings.Trim(name, " "))

	// to lower akan membuat semua karakter menjadi kecil
	fmt.Println(strings.ToLower(name))

	// to Upper akan membuat semua karakter menjadi besar(kapital)
	fmt.Println(strings.ToUpper(name))

	kapitalName := strings.Trim(strings.ToUpper(name), " ")

	// Split untuk memecah string sesuai dengan seperator
	names := strings.Split(kapitalName, " ") //menjadi Slice

	for _, nama := range names {
		fmt.Println(nama)
	}

	// ReplacaAll untuk merubah string yang sama menjadi string yang kita inginkan
	fmt.Println(strings.ReplaceAll(kapitalName, "MAN", "Firman"))

}
