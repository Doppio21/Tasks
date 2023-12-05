package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	Name    string
	Surname string
}

func NewHuman(name string, surname string) Human {
	return Human{
		Name:    name,
		Surname: surname,
	}
}

func (h *Human) Hi() {
	fmt.Printf("Hi! My name is %s %s\n", h.Surname, h.Name)
}

type Action struct {
	Human
	Type string
}

func NewAction(t string, h Human) Action {
	return Action{
		Human: h,
		Type:  t,
	}
}

func (a *Action) Do() {
	fmt.Printf("%s does %s\n", a.Name, a.Type)
}

func main() {
	h := NewHuman("Maria", "Ivanova")
	a := NewAction("washing", h)

	a.Hi()
	a.Do()
}
