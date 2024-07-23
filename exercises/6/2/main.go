package main

import (
	"fmt"
	"time"
)

func goroutine(ch chan struct{}) {
	defer fmt.Println("Горутина завершена")
	for {
		select {
		case <-ch:
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Горутина работает")
		}
	}
}
func main() {
	ch := make(chan struct{})
	go goroutine(ch)
	time.Sleep(2 * time.Second)
	ch <- struct{}{}
	time.Sleep(1 * time.Second)
}
