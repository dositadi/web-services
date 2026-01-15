package main

import "fmt"

type EnglishPerson struct {
	Name string
}

type SpanishPerson struct {
	Name string
}

type Greeter interface {
	Greet() string
}

func (ep EnglishPerson) Greet() string {
	return "Hello, my name is " + ep.Name
}

func (sp SpanishPerson) Greet() string {
	return "Hola, mi nombre es " + sp.Name
}

func main() {
	ep := EnglishPerson{
		Name: "Divine Ositadinma",
	}

	sp := SpanishPerson{
		Name: "Carlos",
	}

	var g Greeter

	g = ep
	fmt.Println(g.Greet())

	g = sp
	fmt.Println(g.Greet())

	fmt.Println()

	SayHello(ep)
	SayHello(sp)
}

func SayHello(greeter Greeter) {
	fmt.Println(greeter.Greet())
}
