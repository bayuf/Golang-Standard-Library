package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("Home/Hello/Hello.go"))
	fmt.Println(path.Base("Home/Hello/Hello.go"))
	fmt.Println(path.Ext("Home/Hello/Hello.go"))
	fmt.Println(path.Join("Home", "Hello", "Hello.go"))

}
