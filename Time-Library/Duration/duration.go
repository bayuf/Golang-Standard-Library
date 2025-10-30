package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := 100 * time.Second
	duration2 := 100 * time.Millisecond
	duration3 := duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("Duration 3 : %d\n", duration3)
}
