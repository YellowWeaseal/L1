package main

import "fmt"

/*
var justString string

	func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = v[:100]
	}

	func main() {
	  someFunc()
	}

Данный фрагмент кода может вызвать панику (panic) из-за попытки обращения к срезу v[:100],
если длина строки v оказывается меньше 100 символов. Это может произойти,
если функция createHugeString возвращает строку, которая короче 100 символов.
Для избежания таких проблем, необходимо проверить, что длина строки v больше или равна 100 символам,
прежде чем пытаться создать срез
*/

var justString string

//пример возможной реализвации createHugeString

func createHugeString(length int) string {
	buffer := make([]byte, length)

	for i := 0; i < length; i++ {
		buffer[i] = 'a'
	}

	return string(buffer)
}

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		// Обработка случая, когда длина строки меньше 100 символов.
		fmt.Println("Длина строки меньше 100 символов")
	}
}

func main() {
	someFunc()
}
