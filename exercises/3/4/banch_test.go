package main

import (
	"math/rand"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum := 0
		arr := make([]int, 0, 60000)
		for i := 0; i < 60000; i++ {
			arr = append(arr, rand.Intn(5000))
		}
		for _, v := range arr {
			sum += v * v
		}
	}
}
