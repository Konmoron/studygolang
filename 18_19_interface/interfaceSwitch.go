package main

import "fmt"

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("String, value = %s\n", i.(string))
	case int:
		fmt.Printf("int, value = %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	findType(56)
	findType("hello world")
	findType(55.5)
}
