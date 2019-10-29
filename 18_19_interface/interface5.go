package main

import "fmt"

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func main() {
	var s interface{} = 55
	assert(s)
	s1 := "hello world"
	assert(s1)
}
