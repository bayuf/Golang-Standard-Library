package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Bayu", "Firmasnyah", "Firman"})
	_ = writer.Write([]string{"Silfi", "Dian", "Putri"})
	_ = writer.Write([]string{"Joko", "Susilo", "Putra"})

	writer.Flush()
}
