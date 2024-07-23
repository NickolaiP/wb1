package main

import (
	"fmt"
	"math"
)

type Points struct {
	point1, point2 [2]float64
}

func (p Points) Distance() float64 {
	fmt.Println(math.Pow(p.point2[1]-p.point1[1], 2))
	return math.Sqrt(math.Pow(p.point2[1]-p.point1[1], 2) + math.Pow(p.point2[0]-p.point1[0], 2))
}
func NewPoints(point1, point2 [2]float64) Points {
	return Points{
		point1: point1,
		point2: point2,
	}
}

func main() {
	p := NewPoints([2]float64{1, 3}, [2]float64{2, 5})
	fmt.Println(p.Distance())
}
