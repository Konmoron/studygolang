package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d year old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a Address) Describe() {
	fmt.Printf("State %s Country %s \n", a.state, a.country)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	var d1 Describer
	p := Person{
		name: "Naveen R",
		age:  54,
	}

	p.Describe()

	d1 = p
	d1.Describe()

	p2 := Person{"James", 22}
	d1 = &p2
	d1.Describe()

	findType(p)

	findType("hello")

	var d2 Describer
	a1 := Address{"beijing", "china"}
	d2 = a1
	d2.Describe()
	d2 = &a1
	d2.Describe()
}
