package main

import "fmt"

func simple(a func(a int, b int) int) {
	fmt.Println(a(60, 7))
}

func main() {
	f := func(a int, b int) int {
		return a + b
	}
	simple(f)
}
