package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range dataChan {
		fmt.Printf("Worker %d получил данные: %d\n", id, data)
	}
	fmt.Printf("Worker %d завершил работу.\n", id)
}

func main() {
	dataChan := make(chan int)
	numWorkers := 3 // Здесь вы можете выбрать количество воркеров

	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg)
	}

	// Обработка сигнала Ctrl+C для завершения программы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		close(dataChan) // Закрыть канал при получении сигнала Ctrl+C
	}()

	// Записываем данные в канал бесконечно
	for i := 1; ; i++ {
		select {
		case dataChan <- i:
		case <-sigChan: // Обработка сигнала Ctrl+C для завершения записи
			break
		}
	}

	wg.Wait() // Ожидаем завершения всех воркеров
}

//Закрытие канала dataChan приводит к тому, что горутины-воркеры, которые ожидают чтения данных из этого канала ,
//замечают, что канал закрыт, и завершают свою работу.
//
//Таким образом, завершение работы всех воркеров происходит путем закрытия канала данных,
