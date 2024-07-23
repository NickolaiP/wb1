package main

import (
	"fmt"
	"sync"
)

func SumSquaring(ch chan int, v int, group *sync.WaitGroup) {
	ch <- v * v
	group.Done()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(arr))
	wg := new(sync.WaitGroup)
	var res = 0
	wg.Add(len(arr))
	for _, v := range arr {
		go SumSquaring(ch, v, wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for elem := range ch {
		res += elem
	}
	fmt.Println(res)
}
