package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func main() {
	p1 := person{
		"John",
		"Smith",
	}

	defer p1.fullName()
	fmt.Printf("Welcome ")
}
