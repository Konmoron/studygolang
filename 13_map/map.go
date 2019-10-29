package main

import "fmt"

func main() {
	personSalary := map[string]int {}
	fmt.Println("map is", personSalary)
	personSalary["mike"] = 9000
	fmt.Println("map is", personSalary)
	for key, value := range personSalary {
		fmt.Println(key, "=", value)
	}
}
