package main

import (
	"fmt"
	"strconv"
)

// references:
// https://stackoverflow.com/a/30544594/6680337

func add(arr *[]string) {
	for i := 0; i < 5; i++ {
		*arr = append(*arr, strconv.Itoa(i))
	}
}

func main() {
	var testArr []string
	add(&testArr)
	fmt.Println(testArr)
}
