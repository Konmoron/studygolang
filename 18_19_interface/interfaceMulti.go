package main

import "fmt"

type SalaryCalculator2 interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperatons interface {
	SalaryCalculator2
	LeaveCalculator
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, e.basicPay+e.pf)
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e1 := Employee{
		firstName: "Naveen",
		lastName: "Ramanathan",
		basicPay: 12000,
		pf: 23423040,
		totalLeaves: 30,
		leavesTaken: 5,
	}

	var d1 SalaryCalculator2 = e1
	d1.DisplaySalary()
	var c1 LeaveCalculator = e1
	fmt.Printf("\nLeaves left = %d", c1.CalculateLeavesLeft())

	var empOp EmployeeOperatons = e1
	empOp.DisplaySalary()
	fmt.Printf("\nLeaves left = %d", empOp.CalculateLeavesLeft())
}
