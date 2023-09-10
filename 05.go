package main

import (
	"fmt"
	"time"
)

func main() {
	dataChan := make(chan int)
	n := 5 // Здесь вы можете установить количество секунд до завершения программы

	// Запускаем горутину для отправки данных в канал
	go func() {
		for i := 1; ; i++ {
			dataChan <- i
		}
	}()

	// Запускаем горутину для чтения данных из канала
	go func() {
		for {
			data := <-dataChan
			fmt.Printf("Прочитано из канала: %d\n", data)
		}
	}()

	// Устанавливаем таймер на завершение программы через N секунд
	timer := time.NewTimer(time.Duration(n) * time.Second)
	<-timer.C // Ожидаем срабатывания таймера

	// Завершаем программу
	fmt.Printf("Программа завершена по истечении %d секунд.", n)
}
