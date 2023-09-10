package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}

func main() {

	fmt.Println("Начало программы")

	seconds := 3
	Sleep(seconds)

	fmt.Printf("Прошло %d секунды", seconds)
}
