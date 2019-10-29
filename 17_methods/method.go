package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

// display func
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	epm1 := Employee{
		name:     "Sam Adolf",
		salary:   50000,
		currency: "$",
	}

	epm1.displaySalary()

	r := Rectangle{
		length: 10,
		width:  20,
	}

	fmt.Printf("\nArea of Rectangle is %d", r.Area())

	c := Circle{
		radius: 20,
	}

	fmt.Printf("\nArea of Circle is %f", c.Area())
}
