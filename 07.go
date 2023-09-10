package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[string]int)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	numWorkers := 5
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			defer wg.Done()

			// Запись данных в map
			key := fmt.Sprintf("key%d", id)
			value := id * 10

			mutex.Lock()
			data[key] = value
			mutex.Unlock()

			fmt.Printf("Worker %d записал в map: %s -> %d\n", id, key, value)
		}(i)
	}

	wg.Wait()

	// Вывод данных из map
	fmt.Println("Содержимое map:")
	for key, value := range data {
		fmt.Printf("%s -> %d\n", key, value)
	}
}
