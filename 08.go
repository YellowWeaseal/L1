package main

import (
	"fmt"
)

func setBit(n int64, pos uint, val bool) int64 {
	if val {
		// Устанавливаем i-й бит в 1
		return n | (1 << pos)
	} else {
		// Устанавливаем i-й бит в 0
		return n &^ (1 << pos)
	}
}

func main() {
	var num int64 = 1847591 // Пример исходного числа
	bitPos := uint(6)       // Номер изменяемого бита

	// Устанавливаем бит в 1
	result := setBit(num, bitPos, true)
	fmt.Printf("Установлен %d-й бит в 1: %d\n", bitPos, result)

	// Устанавливаем бит в 0
	result = setBit(num, bitPos, false)
	fmt.Printf("Установлен %d-й бит в 0: %d\n", bitPos, result)
}
