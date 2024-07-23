package main

import (
	"fmt"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(DeleteElement([]int{1, 2, 3, 4, 5, 6}, 2))
	}
}
