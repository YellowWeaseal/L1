package main

import "fmt"

func main() {
	a := 7
	b := 15

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	// Меняем местами a и b
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}
