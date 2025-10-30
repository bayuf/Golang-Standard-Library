package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var datas *ring.Ring = ring.New(5)

	for i := 0; i < datas.Len(); i++ {
		datas.Value = "Value " + strconv.Itoa(i)
		datas = datas.Next()
	}

	datas.Do(func(value any) {
		fmt.Println(value)
	})
}
