package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(DeleteElement([]int{1, 2, 3, 4, 5, 6}, 2))
	fmt.Println(DeleteElement([]string{"dog", "cat", "tiger", "cucumber"}, 3))
}

func DeleteElement[T any](s []T, index int) []T {
	if index < 0 || index >= len(s) {
		log.Fatal("Index out of range")
	}
	buffer := make([]T, 0, len(s)-1)
	buffer = append(buffer, s[:index]...)
	buffer = append(buffer, s[index+1:]...)
	return buffer
}
