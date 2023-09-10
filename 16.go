package main

import "fmt"

//1 вариант
/*type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func main() {
	arr := []int{64, 25, 12, 22, 11}
	intSlice := IntSlice(arr)

	// Вызываем функцию sort.Sort, передавая intSlice и ее методы интерфейса
	sort.Sort(intSlice)

	fmt.Println("Отсортированный массив:", intSlice)
}*/

// 2 вариант

/*func main() {
	arr := []int{64, 25, 12, 22, 11}

	sort.Ints(arr)

	fmt.Println("Отсортированный массив:", arr)
}*/

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, num := range arr[1:] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Исходный массив:", arr)
	arr = quickSort(arr)
	fmt.Println("Отсортированный массив:", arr)
}
