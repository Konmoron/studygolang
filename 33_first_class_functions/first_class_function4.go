package main

import "fmt"

func simple1() func(a int, b int) int {
	f := func(a int, b int) int {
		return a + b
	}

	return f
}

func main() {
	s := simple1()
	fmt.Println(s(60, 8))
}
