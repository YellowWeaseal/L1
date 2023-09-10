package main

import (
	"fmt"
	"sync"
)

func main() {
	inputChan := make(chan int)
	outputChan := make(chan int)
	var wg sync.WaitGroup

	// Горутина для чтения чисел из массива и отправки их в inputChan
	wg.Add(1)
	go func() {
		defer close(inputChan)
		defer wg.Done()

		numbers := []int{1, 2, 3, 4, 5}

		for _, num := range numbers {
			inputChan <- num
		}
	}()

	// Горутина для умножения чисел на 2 и отправки результатов в outputChan
	wg.Add(1)
	go func() {
		defer close(outputChan)
		defer wg.Done()

		for num := range inputChan {
			result := num * 2
			outputChan <- result
		}
	}()

	// Горутина для вывода результатов в stdout
	wg.Add(1)
	go func() {
		defer wg.Done()

		for result := range outputChan {
			fmt.Println(result)
		}
	}()

	// Ожидание завершения всех горутин
	wg.Wait()
}
