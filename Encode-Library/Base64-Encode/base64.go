package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	name := "Bayu Firmansyah"

	encoded := base64.StdEncoding.EncodeToString([]byte(name)) // String ke byte

	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}
