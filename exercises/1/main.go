package main

import "fmt"

type human struct {
	name string
	age  int
}

func (h human) sayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years old\n", h.name, h.age)
}

type action struct {
	human
}

func main() {
	h := human{name: "Alice", age: 30}
	a := action{human: h}

	a.sayHello()
}
