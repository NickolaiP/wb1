package main

import "fmt"

func determineType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Type: int, Value: %d\n", v)
	case string:
		fmt.Printf("Type: string, Value: %s\n", v)
	case bool:
		fmt.Printf("Type: bool, Value: %t\n", v)
	case float64:
		fmt.Printf("Type: float64, Value: %f\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	var i interface{}

	i = 42
	determineType(i)

	i = "Hello, World!"
	determineType(i)

	i = true
	determineType(i)

	i = 10.0
	determineType(i)

}
