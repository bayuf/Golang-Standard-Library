package main

import (
	"fmt"
	"math"
)

func main() {
	angka := 10.9

	// Round Untuk membulatkan angka ke atas atau ke bawah, sesuai dengan yang paling dekat
	fmt.Println(math.Round(angka)) //11

	// Floor untuk membulatkan angka ke bawah
	fmt.Println(math.Floor(angka)) //10

	// Ceil untuk membulatkan angka ke atas
	fmt.Println(math.Ceil(angka)) //11

	// Max untuk mencari angka terbesar dari 2 angka
	fmt.Println(math.Max(10, 15)) //15

	// Min untuk mencari angka terkecil dari 2 angka
	fmt.Println(math.Min(10, 15)) //10

}
