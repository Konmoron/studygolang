package main

import (
	"fmt"
	"studygolang_com/structs/computer"
)

func main() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 40
	fmt.Println("spec", spec)
}
