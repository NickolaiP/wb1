package main

import (
	"sync"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numbers := []int{2, 4, 6, 8, 10}
		resultChannel := make(chan int)
		var wg sync.WaitGroup

		for _, num := range numbers {
			wg.Add(1)
			go func(n int) {
				defer wg.Done()
				calculateSquare(n, resultChannel)
			}(num)
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
