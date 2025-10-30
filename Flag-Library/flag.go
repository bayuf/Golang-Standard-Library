package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("hostname", "localhost", "Put your database host")
	userName := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	port := flag.Int("port", 0, "Put your database port")

	flag.Parse()

	fmt.Println("Hostname :", *host)
	fmt.Println("Username :", *userName)
	fmt.Println("Password :", *password)
	fmt.Println("Port :", *port)

}
