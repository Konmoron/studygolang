package main

import "fmt"

func main() {
	personSalary := map[string]int{}
	fmt.Println("map is", personSalary)
	personSalary["mike"] = 9000
	fmt.Println("map is", personSalary)
	for key, value := range personSalary {
		fmt.Println(key, "=", value)
	}

	testMap := map[string]string{"A": "A", "B": "B", "C": "C"}
	for k, v := range testMap {
		fmt.Println("key:", k, ", value", v)
	}
}
