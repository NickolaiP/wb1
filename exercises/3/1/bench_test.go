package main

import (
	"sync"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{2, 4, 6, 8, 10}
		wg := new(sync.WaitGroup)
		mu := new(sync.Mutex)
		var sum int

		for _, i := range arr {
			wg.Add(1)
			go SumSquaring(i, &sum, wg, mu)
		}
		wg.Wait()
	}
}
