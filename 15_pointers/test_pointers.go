package main

import "fmt"

func main() {
	b := "hello world"
	var a *string
	if a == nil {
		fmt.Println("a is nil")
		a = &b
		fmt.Println("a after initialization is", a)
		fmt.Println("value of a is", *a)
		*a += "!"
	}
	c := &b
	fmt.Printf("Type of a is %T\n", c)
	fmt.Println("address of b is", c)
	fmt.Println("value of c is", *c)
}
