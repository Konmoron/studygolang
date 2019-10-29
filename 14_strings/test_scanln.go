package main

import (
	"fmt"
)

func main() {
	var firstName, lastName string
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
}
