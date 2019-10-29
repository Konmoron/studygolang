package main

import (
	"fmt"
	"math"
)

type Employee1 struct {
	name, currency string
	salary         int
}

func (e Employee1) changeName(newName string) {
	e.name = newName
}

func (e *Employee1) changeSalary(newSalary int) {
	e.salary = newSalary
}

func (e Employee1) displaySalary() {
	fmt.Printf("\nSalary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func displaySalaryFunc(e Employee1) {
	fmt.Printf("\nSalary of %s is %s%d\n", e.name, e.currency, e.salary)
}

type Rectangle struct {
	length, width int
}

func perimeter(r *Rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

func (r *Rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
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

type address struct {
	city, state string
}

func (a address) fullAddress() {
	fmt.Printf("\nFull Address: %s, %s\n", a.city, a.state)
}

type person struct {
	firstName, lastName string
	address
}

func main() {
	emp1 := Employee1{"Sam", "$", 140000}
	emp1.displaySalary()
	displaySalaryFunc(emp1)

	fmt.Printf("Employee name before change: %s\n", emp1.name)
	emp1.changeName("Tom")
	fmt.Printf("Employee name after change: %s\n", emp1.name)

	fmt.Printf("Employee salary before change: %d\n", emp1.salary)
	emp1.changeSalary(1500000)
	fmt.Printf("Employee salary after change: %d\n", emp1.salary)

	r1 := Rectangle{15, 15}
	fmt.Printf("\nArea of rectangle is %d\n", r1.Area())

	r1.perimeter()
	r1P := &r1
	r1P.perimeter()
	perimeter(r1P)
	//perimeter(r1)

	c1 := Circle{3}
	fmt.Printf("\nArea of circle is %f", c1.Area())

	p1 := person{"tom", "sam", address{"Los Angeles", "California",},}
	p1.fullAddress()
}
