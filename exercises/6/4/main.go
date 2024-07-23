package main

import (
	"fmt"
	"time"
)

func goroutine(ch chan string) {
	defer fmt.Println("Горутина завершена")
	for v := range ch {
		time.Sleep(time.Second)
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan string)
	go goroutine(ch)
	go func(ch chan string) {
		ch <- "Горутина работает"
	}(ch)
	time.Sleep(2 * time.Second)
	close(ch)
	time.Sleep(1 * time.Second)
}
