package main

import (
	"fmt"
	"sync"
)

func SumSquaring(i int, sum *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	*sum += i * i
	mu.Unlock()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	var sum int

	for _, i := range arr {
		wg.Add(1)
		go SumSquaring(i, &sum, wg, mu)
	}
	wg.Wait()
	fmt.Println(sum)
}
