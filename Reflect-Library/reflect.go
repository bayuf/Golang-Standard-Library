package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` //struct tag
}

type Person struct {
	Name string `required:"true" max:"10"`
	Age  int    `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name :", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		fmt.Println(structField.Name, "dengan tipe ", structField.Type)
		fmt.Println(structField.Tag.Get("required")) // cek apakah ada struct tag required
	}
}

func isValid(datas any) (result bool) {
	result = true
	// mendapatkan type data
	t := reflect.TypeOf(datas)
	// mendapatkan field
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(datas).Field(i).Interface()
			result = data != ""
			// akan return kalau menemukan data kosong di field yang ada tag "required"
			if result == false {
				return result
			}
		}
	}

	return result
}

func main() {
	// readField(Sample{"Bayu"})
	// readField(Person{"Bayu", 25})

	person := Person{
		Name: "Bayu",
		Age:  25,
	}

	fmt.Println(isValid(person))
}
