package main

import (
	"fmt"
)

func reverseString(input string) string {
	// Разбиваем строку на юникодные символы
	runes := []rune(input)
	length := len(runes)

	// Создаем новый слайс для перевернутых символов
	reversed := make([]rune, length)

	// Переворачиваем символы
	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		reversed[i] = runes[j]
	}

	return string(reversed)
}

func main() {
	input := "тыпок мотопот доп"
	reversed := reverseString(input)
	fmt.Println(reversed)
}
