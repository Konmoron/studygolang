package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permenant struct {
	empId int
	basicpay int
	pf int
}

type Contract struct {
	empId int
	basicpay int
}

func (p Permenant) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s{
		expense += v.CalculateSalary()
	}
	fmt.Printf("Total expense is $%d", expense)
}

func main() {
	pemp1 := Permenant{1, 10000, 20000}
	pemp2 := Permenant{2, 120000, 400000}
	cemp1 := Contract{1, 150000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}
