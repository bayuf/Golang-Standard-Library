package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvData := "bayu, firmansyah, firman\n" +
		"silfi, Dian, Putri\n" +
		"Joko, andi, pratama"

	reader := csv.NewReader(strings.NewReader(csvData))

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
