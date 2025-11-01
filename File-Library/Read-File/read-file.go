package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)

	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	message := ""

	for {
		line, _, err := reader.ReadLine()
		message += string(line) + "\n"
		if err == io.EOF {
			break
		}

	}

	return message, nil
}

func main() {
	result, err := readFile("teks.txt")
	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(result)
	}
}
