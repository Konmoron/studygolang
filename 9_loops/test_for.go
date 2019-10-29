package main

import "fmt"

func main() {
	num := 10
	for i := 0; i < 10; i++{
		fmt.Println("num is", num - i)
	}
	i1 := 0
	for i1 < 10 {
		fmt.Println(i1)
		i1 += 2
	}
	for {
		fmt.Println("hello world")
	}
}
