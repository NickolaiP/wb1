package main

import (
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 0, 60000)
		for i := 0; i < 60000; i++ {
			arr = append(arr, rand.Intn(5000))
		}
		core := 6
		buf := make([]int, len(arr)/core)
		step := len(arr) / core
		resultChannel := make(chan int, len(arr))
		wg := new(sync.WaitGroup)
		for i := 0; i < core; i++ {
			wg.Add(1)
			copy(buf, arr[i*step:(i+1)*step])
			go calculateSquare(buf, resultChannel, wg)
		}

		go func() {
			wg.Wait()
			close(resultChannel)
		}()

		sum := 0
		for square := range resultChannel {
			sum += square
		}
	}
}
