package main

import (
	"fmt"
)

type Person struct {
	name string
}

func main() {
	var p *Person
	p = createPerson("John")
	fmt.Println("Person:", p.name)
}

func createPerson(name string) *Person {
	person := Person{name: name}
	return &person
}