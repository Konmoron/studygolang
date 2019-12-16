package main

import (
	"fmt"
	"studygolang/16_structures/structs/computer"
)

func main() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 40
	fmt.Println("spec", spec)
}
