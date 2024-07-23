package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	arr := make([]int, 0, 1200)
	for i := 0; i < 1200; i++ {
		arr = append(arr, rand.Intn(5000))
	}
	core := 6
	b := make([]int, len(arr)/core)
	step := len(arr) / core
	resultChannel := make(chan int, len(arr))
	wg := new(sync.WaitGroup)
	for i := 0; i < core; i++ {
		wg.Add(1)
		copy(b, arr[i*step:(i+1)*step])
		go calculateSquare(b, resultChannel, wg)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	sum := 0
	for square := range resultChannel {
		sum += square
	}
	fmt.Println(sum)
}

func calculateSquare(nums []int, resultChannel chan int, wg *sync.WaitGroup) {
	for _, num := range nums {
		resultChannel <- num * num
	}
	wg.Done()
}
