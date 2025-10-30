package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Bayu")
	data.PushBack("Firmansyah")
	data.PushBack(25)

	front := data.Front()
	fmt.Println(front.Value) //Bayu

	next := front.Next()
	fmt.Println(next.Value) //Firmansyah

	next = next.Next()
	fmt.Println(next.Value) //25

	// Menggunakan Iterasi For
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
