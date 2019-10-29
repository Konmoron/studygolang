package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("channel a type is %T\n", a)
	}

	b := make(chan string)
	fmt.Printf("channel b type is %T", b)
}
