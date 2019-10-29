package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("hello world from first class function!")
	}
	a()
	fmt.Printf("%T\n", a)

	func() {
		fmt.Println("hello world from first class function 2!")
	}()

	func(s string) {
		fmt.Println("hello", s)
	}("Tom")
}
