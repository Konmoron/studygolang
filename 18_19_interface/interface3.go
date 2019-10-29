package main

import "fmt"

type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describe(t Test) {
	fmt.Printf("18_19_interface type %T value %v\n", t, t)
}

func main() {
	var t Test
	f := MyFloat(34.4)
	t = f
	t.Tester()
	describe(t)
}