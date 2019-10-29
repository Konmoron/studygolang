package main

import (
	"fmt"
)

func main() {
	var a []string
	fmt.Printf("a %p\n", a)

	fmt.Printf("a %p\n", a)
	fmt.Println("Hello, playground")
}