package main

import "fmt"

//1 вариант

/*func binarySearch(arr []int, target int) int {
	index := sort.SearchInts(arr, target)
	if index < len(arr) && arr[index] == target {
		return index
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7

	fmt.Printf("Исходный массив: %v\n", arr)
	fmt.Printf("Искомый элемент: %d\n", target)

	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}*/

// 2 вариант
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // Если элемент не найден
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7

	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден по индексу %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
