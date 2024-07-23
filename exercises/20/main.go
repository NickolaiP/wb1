package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	arr := strings.Split(s, " ")
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return strings.Join(arr, " ")
}

func main() {
	fmt.Println(Reverse("snow dog sun"))
}
