package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary 		int
}

type Address struct {
	city, state string
}

type Person struct {
	name string
	age int
	address Address
}

type PersonPromoted struct {
	name string
	age int
	Address
}

func main() {
	emp1 := Employee{
		firstName: "Sam",
		age: 25,
		salary: 190000,
		lastName: "Anderson",
	}

	emp2 := Employee{"Thomas", "Paul", 25, 100000}

	emp3 := struct {
		firstName, lastName string
		age, salary			int
	}{"And", "Nikola", 24, 23000}

	var emp4 Employee
	emp4.firstName = "Jack"
	emp4.lastName = "Adams"

	emp5 := Employee{
		firstName: "John",
		lastName: "Paul",
	}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
	fmt.Println("Employee 3", emp3)
	fmt.Println("Employee 4", emp4)
	fmt.Println("Employee 5", emp5)

	emp6 := Employee{"Sam", "And", 25, 140004}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Println("Salary:", emp6.salary)

	var person1 Person
	person1.age = 50
	person1.name = "Naveen"
	person1.address = Address{"Chicago", "Illinois"}
	fmt.Println("person 1", person1)
	fmt.Println("City", person1.address.city)

	var personPromoted PersonPromoted
	personPromoted.name = "Nave"
	personPromoted.age = 44
	personPromoted.Address = Address{"Chicago", "Illinois"}
	fmt.Println("City", personPromoted.city)

}
