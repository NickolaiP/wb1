package main

import (
	"context"
	"fmt"
	"time"
)

func goroutine(ctx context.Context) {
	defer fmt.Println("Горутина завершена")
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Горутина работает")
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go goroutine(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
