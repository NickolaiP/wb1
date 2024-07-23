package main

import (
	"fmt"
	"reflect"
)

func determineType(i interface{}) {
	switch reflect.TypeOf(i).Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Println("channel")
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	var i interface{}
	i = 42
	determineType(i)
	i = "hello"
	determineType(i)
	i = true
	determineType(i)
	i = make(chan int)
	determineType(i)
}
