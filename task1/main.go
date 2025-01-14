package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SayHello() {
	fmt.Printf("Hello wb team, my name is %s, I'm %d years old\n", h.Name, h.Age)
}

type Action struct {
	Human 
}

func (a *Action) DoAction() {
	fmt.Println("Call from Action method")
}

func main() {
	/*Дана структура Human (с произвольным набором полей и методов). 
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
	*/
	action := Action{
		Human: Human{
			Name: "Bogdan",
			Age:  24,
		},
	}
	action.SayHello() //call Human struct method
	action.DoAction() //call Action struct method
}
