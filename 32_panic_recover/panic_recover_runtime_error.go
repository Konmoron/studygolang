package main

import (
	"fmt"
	"runtime/debug"
)

func recoverName2() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
		debug.PrintStack()
	}
}

func c() {
	defer recoverName2()
	n := []int{1, 3, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func main() {
	c()
	fmt.Println("normally return from main")
	//debug.PrintStack()
}
