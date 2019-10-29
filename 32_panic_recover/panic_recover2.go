package main

import (
	"fmt"
	"time"
)

func recoverName1() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
	}
}

func a() {
	defer recoverName1()
	fmt.Println("inside a")
	//go b()
	b()
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("inside b")
	panic("oh! panic from b")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}
