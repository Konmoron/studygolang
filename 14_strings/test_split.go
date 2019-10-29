package main

import (
	"fmt"
	"strings"
)

func printSlice(s string) {
	arr := strings.Split(s, ",")
	fmt.Printf("s: '%s', len of arr: %d, arr: '%s'\n", s, len(arr), arr)
	for i, v := range arr {
		//fmt.Println("i:", i, ", v:'", v, "'")
		fmt.Printf("i: %d, v: '%s'\n", i, v)
	}
	fmt.Println()
}

func main() {
	s1 := ""
	printSlice(s1)
	s2 := "a,b"
	printSlice(s2)
	s3 := "a, b"
	printSlice(s3)
	s4 := ","
	printSlice(s4)
	s5 := "a,"
	printSlice(s5)
	s6 := ",b"
	printSlice(s6)
}
