package main

import "fmt"

func main() {
	// Исходный срез
	slice := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	indexToRemove := 1

	// Проверка на валидность индекса
	if indexToRemove >= 0 && indexToRemove < len(slice) {
		// Удаление элемента по индексу
		slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
		fmt.Println(slice)
	} else {
		fmt.Println("Неверный индекс для удаления")
	}
}
