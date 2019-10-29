package main

import "fmt"

func printA(a int) {
	fmt.Println("Value of a in deferred function is", a)
}

func main() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("Value of a in main function is", a)
}
