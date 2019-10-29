package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 12.5, 16.4
	d := math.Min(a, b)
	fmt.Println("min num is", d)
	fmt.Printf("type of a is %T\n", a)
	c := 10
	sum := float64(c) + a
	fmt.Println("sum is", sum)
}
