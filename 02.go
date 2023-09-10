package main

import (
	"fmt"
	"sync"
)

func Squared(numbers []int) {
	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			square := n * n
			fmt.Printf("Квадрат числа %d равен %d\n", n, square)
		}(num)
	}
	wg.Wait()
}

func main() {
	numbers := []int{2, 4, 6, 8}
	Squared(numbers)
}
