package main

import (
	"sync"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr = []int{2, 4, 6, 8, 10}
		var ch = make(chan int, len(arr))
		var wg = new(sync.WaitGroup)
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
	}
}
