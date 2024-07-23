package main

import (
	"fmt"
	"sync"
)

func Squaring(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i * i)
}
func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := new(sync.WaitGroup)
	for _, i := range arr {
		wg.Add(1)
		go Squaring(i, wg)
	}
	wg.Wait()
}
