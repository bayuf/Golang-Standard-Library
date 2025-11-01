package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("Home/Hello/main.go"))
	fmt.Println(filepath.Base("Home/Hello/main.go"))
	fmt.Println(filepath.Ext("Home/Hello/main.go"))
	fmt.Println(filepath.IsAbs("Home/Hello/main.go"))
	fmt.Println(filepath.IsLocal("Home/Hello/main.go"))
	fmt.Println(filepath.Join("Home", "Hello", "main.go"))

}
