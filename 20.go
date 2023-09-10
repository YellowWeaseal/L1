package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	// Разбиваем строку на слова
	words := strings.Fields(input)

	// Создаем новый слайс для перевернутых слов
	reversedWords := make([]string, len(words))

	// Переворачиваем слова в обратном порядке и помещаем их в новый слайс
	for i, j := 0, len(words)-1; i < len(words); i, j = i+1, j-1 {
		reversedWords[i] = words[j]
	}

	// Объединяем перевернутые слова в строку с пробелами между ними
	result := strings.Join(reversedWords, " ")

	return result
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println(reversed)
}
