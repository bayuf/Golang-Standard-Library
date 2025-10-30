package main

import (
	"fmt"
	"os"
)

func main() {
	// Args digunakan untuk menangkap argumen pada command line
	args := os.Args

	for _, v := range args {
		fmt.Println(v)
	}

	hostName, err := os.Hostname()

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(hostName)
	}
}
