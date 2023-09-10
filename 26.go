package main

import (
	"fmt"
	"strings"
)

func areAllCharactersUnique(s string) bool {
	// Приводим строку к нижнему регистру
	s = strings.ToLower(s)

	// Создаем карту для отслеживания уникальных символов
	characters := make(map[rune]bool)

	// Перебираем символы строки
	for _, char := range s {
		if characters[char] {
			// Если символ уже был встречен, то он не уникален
			return false
		}
		characters[char] = true
	}

	// Если все символы уникальны
	return true
}

func main() {
	stringsToCheck := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, s := range stringsToCheck {
		if areAllCharactersUnique(s) {
			fmt.Printf("%s — true\n", s)
		} else {
			fmt.Printf("%s — false\n", s)
		}
	}
}
