package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

type Worker struct {
	ctx context.Context
	wg  *sync.WaitGroup
	ch  <-chan interface{}
}

func (w Worker) Run() {

	defer w.wg.Done()
	for {
		select {
		case value, ok := <-w.ch:
			if ok {
				fmt.Println(value)
			}
		case <-w.ctx.Done():
			return
		}
	}
}
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <N>")
		os.Exit(1)
	}

	N, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Invalid argument for N")
		os.Exit(0)
	}

	signalShurtDown := make(chan os.Signal, 1)
	signal.Notify(signalShurtDown, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	mainCh := make(chan interface{})
	wg := new(sync.WaitGroup)
	worker := Worker{ctx, wg, mainCh}

	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker.Run()
	}

	for {
		select {
		case <-signalShurtDown:
			close(mainCh)
			cancel()
			wg.Wait()
			return
		default:
			mainCh <- rand.Intn(100)
		}
	}
}
