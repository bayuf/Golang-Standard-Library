package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`b([a-z][a-z])u`)

	fmt.Println(regex.MatchString("bayu"))
	fmt.Println(regex.MatchString("bazu"))
	fmt.Println(regex.MatchString("bAYu"))

	fmt.Println(regex.FindAllString("bayu Bayu hayu baya brtu bweu", 10))

}
