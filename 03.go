package main

import (
	"fmt"
	"sync"
)

func SumSquared(numbers []int) int {
	result := make(chan int)
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			square := n * n
			result <- square
		}(num)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	sum := 0
	for sqr := range result {
		sum += sqr
	}
	return sum
}
func main() {
	numbers := []int{2, 4, 6, 8, 10}
	result := SumSquared(numbers)

	fmt.Printf("Сумма квадратов чисел: %d\n", result)
}
