package main

import "fmt"

func appendStr() func(s string) string {
	t := "hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func main() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("Tom"))
	fmt.Println(b("Everyone"))
	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}
