package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2008, time.December, 17, 0, 0, 0, 0, time.UTC) //2008-12-17 00:00:00
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2000-05-19 22:30:00"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(valueTime.Local())
	}

	fmt.Println("Year :", valueTime.Year())
	fmt.Println("Month :", valueTime.Month())
	fmt.Println("Day :", valueTime.Day())
	fmt.Println("Hour :", valueTime.Hour())

}
