package main

import (
	"fmt"
	"time"
)

var stopFlag bool

func goroutine() {
	defer fmt.Println("Горутина завершена")
	for {
		fmt.Println("Горутина работает")
		time.Sleep(1 * time.Second)
		if stopFlag {
			break
		}
	}
}

func main() {
	go goroutine()
	time.Sleep(5 * time.Second)
	stopFlag = true
	time.Sleep(1 * time.Second)
}
