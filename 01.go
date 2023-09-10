package main

import "fmt"

// Создаем структуру Action  определяем для нее методы
type Action struct {
	Sleeping bool
	Eating   bool
	Working  bool
}

// Создаем структуру Human которая наследует методы Action
type Human struct {
	Name    string
	Age     int
	Address string
	Action
}

func (a *Action) Sleep() {
	a.Sleeping = true
}

func (a *Action) Eat() {
	a.Eating = true
}
func (a *Action) Work() {
	a.Working = false
}
func main() {
	// Создаем объект Human
	h := Human{
		Name:    "Ivan",
		Age:     23,
		Address: "Moscow",
	}
	// Вызываем методы  Action через структуру Human
	h.Eat()
	h.Work()
	h.Sleep()

	result := fmt.Sprintf("Name:%s, Age:%d, Address:%s, Eating:%t, Working:%t, Sleeping:%t",
		h.Name, h.Age, h.Address, h.Eating, h.Action.Working, h.Sleeping)
	fmt.Println(result)
}
