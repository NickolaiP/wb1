package main

import "fmt"

func main() {
	a := []int{1, 2, 2, 3, 5}
	b := []int{1, 2, 5}
	fmt.Println(intersection(a, b))

}

func intersection(arr1, arr2 []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	for _, v := range arr1 {
		_, ok := m[v]
		if ok {
			continue
		} else {
			m[v] = 1
		}
	}

	for _, v := range arr2 {
		_, ok := m[v]
		if ok && m[v] == 1 {
			res = append(res, v)
			m[v] = 2
		} else {
			continue
		}
	}

	return res
}
