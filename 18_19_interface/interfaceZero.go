package main

import "fmt"

type Describer2 interface {
	Describe()
}

func main() {
	var d1 Describer2

	if d1 == nil {
		fmt.Printf("d1 is nil, type is %T, value is %v\n", d1, d1)
	}
}
