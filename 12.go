package main

import (
	"fmt"
)

func createSet(strings []string) map[string]bool {
	set := make(map[string]bool)

	for _, str := range strings {
		set[str] = true
	}

	return set
}

func main() {
	// Последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем собственное множество
	mySet := createSet(strings)

	// Выводим уникальные элементы множества
	fmt.Println("Собственное множество:")
	for str := range mySet {
		fmt.Println(str)
	}
}
