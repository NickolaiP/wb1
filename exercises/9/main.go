package main

import (
	"fmt"
)

func main() {
	input := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go stage1(input, ch1)
	go stage2(ch1, ch2)

	for result := range ch2 {
		fmt.Println(result)
	}
}

func stage1(input []int, ch1 chan int) {
	for _, num := range input {
		ch1 <- num
	}
	close(ch1)
}

func stage2(ch1 chan int, ch2 chan int) {
	for num := range ch1 {
		result := num * 2
		ch2 <- result
	}
	close(ch2)
}
