package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

type Reciver struct {
	ctx context.Context
	wg  *sync.WaitGroup
	ch  chan interface{}
}

func (r Reciver) Run() {
	defer r.wg.Done()
	for {
		select {
		case <-r.ctx.Done():
			return
		default:
			fmt.Println(<-r.ch)
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

	wg := new(sync.WaitGroup)
	ch := make(chan interface{})
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	defer cancel()
	rec := Reciver{ctx, wg, ch}
	wg.Add(1)
	go rec.Run()

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return
		default:
			time.Sleep(200 * time.Millisecond)
			ch <- rand.Intn(100)
		}
	}
}
