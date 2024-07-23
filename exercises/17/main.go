package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	sort.Ints(arr)
	fmt.Println(binarySearch(arr, 8))
}

func binarySearch(arr []int, target int) int {
	leftPointer := 0
	rightPointer := len(arr) - 1
	for leftPointer <= rightPointer {
		var midPointer = (leftPointer + rightPointer) / 2
		var midValue = arr[midPointer]
		if midValue == target {
			return midPointer
		} else if midValue < target {
			leftPointer = midPointer + 1
		} else if midValue > target {
			rightPointer = midPointer - 1
		}
	}
	return -1
}
