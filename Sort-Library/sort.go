package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// Panjang Slice
func (s UserSlice) Len() int {
	return len(s)
}

// Perbandingan umur user
func (s UserSlice) Less(i, j int) bool {
	return s[i].Age > s[j].Age
}

// Pemindahan Posisi
func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Bayu", 25},
		{"Firman", 28},
		{"Budi", 19},
		{"Joko", 17},
	}

	// Proses melakukan sorting menggunakan interface Sort pada package sort
	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
