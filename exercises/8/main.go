package main

import "fmt"

func setBit(num int64, bitPosition int, bitValue int) int64 {
	mask := int64(1 << uint(bitPosition))
	if bitValue == 1 {
		num |= mask
	} else {
		num &= ^mask
	}
	return num
}

func main() {
	var num int64 = 4 // Пример исходного числа
	bitPosition := 1  // Устанавливаемый бит
	bitValue := 1     // Значение, 1 или 0

	num = setBit(num, bitPosition, bitValue)

	fmt.Printf("Результат после установки бита: %d\n", num)
}
