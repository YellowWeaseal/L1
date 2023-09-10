package main

import (
	"fmt"
)

func getType(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	default:
		return "unknown"
	}
}

func main() {
	var a interface{} = 42
	var b interface{} = "Hello"
	var c interface{} = true
	var d interface{} = make(chan int)

	fmt.Printf("a is of type: %s\n", getType(a))
	fmt.Printf("b is of type: %s\n", getType(b))
	fmt.Printf("c is of type: %s\n", getType(c))
	fmt.Printf("d is of type: %s\n", getType(d))
}
