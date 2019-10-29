package main

import "fmt"

func main() {
	name := "Naveen"
	fmt.Println("Original string is", string(name))
	fmt.Printf("Reversed string is ")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
