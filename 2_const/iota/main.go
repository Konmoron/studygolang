package main

import "fmt"

type Level int8

const (
	A = iota
	B
	C = iota - 2
	D
	E
	F
)

func main() {
	fmt.Println("A:", A)
	fmt.Println("B:", B)
	fmt.Println("C:", C)
	fmt.Println("D:", D)
	fmt.Println("D:", E)
	fmt.Println("F:", F)
}
