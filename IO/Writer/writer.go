package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	_, _ = writer.WriteString("Hallo Dunia\n")
	_, _ = writer.WriteString("Go Programming\n")

	writer.Flush()
}
