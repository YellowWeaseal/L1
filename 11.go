package main

import (
	"fmt"
)

func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	// Проходим по элементам первого множества
	for elem := range set1 {
		// Если элемент присутствует и во втором множестве, добавляем его в результат
		if set2[elem] {
			result[elem] = true
		}
	}

	return result
}

func main() {
	// Примеры двух неупорядоченных множеств
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	// Выполняем пересечение
	result := intersection(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение множеств:")
	for elem := range result {
		fmt.Println(elem)
	}
}
