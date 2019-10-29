package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAmount int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netIncome = 0
	for _, v := range ic {
		fmt.Printf("Income from project %s = $%d\n", v.source(), v.calculate())
		netIncome += v.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netIncome)
}

func main() {
	project1 := FixedBilling{"project1", 5234324}
	project2 := TimeAndMaterial{"project2", 234, 23452}
	project3 := FixedBilling{"project3", 234234}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}