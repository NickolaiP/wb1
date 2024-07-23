package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sort.Float64s(arr)
	fmt.Println(groupTemp(arr))

}

func groupTemp(arr []float64) map[int][]float64 {
	if len(arr) == 0 {
		return nil
	}
	group := make(map[int][]float64)
	key := int(arr[0]) / 10 * 10
	arrres := make([]float64, 0, len(arr))
	for _, v := range arr {
		if v-float64(key) < 10 {
			arrres = append(arrres, v)
		} else {
			group[key] = arrres
			key = int(v) / 10 * 10
			arrres = nil
			arrres = append(arrres, v)
		}
	}
	group[key] = arrres
	return group
}
