package main

import (
	"fmt"
	"strconv"
)

func main() {

	// String ke Boolean
	kondisi, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(kondisi)
	}

	// String ke Integer
	angka, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(angka)
	}

	// Integer ke String
	angkaStr := strconv.FormatInt(9999, 2)
	fmt.Println(angkaStr)

	angkaStr2 := strconv.Itoa(100)
	fmt.Println(angkaStr2)

	// Float ke String
	angkaFloat, err := strconv.ParseFloat("255.5", 32)
	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(angkaFloat)
	}

}
