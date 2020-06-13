package main

import (
	"fmt"
	"strconv"
)

func add(c chan string) {
	for i := 0; i < 5; i++ {
		c <- strconv.Itoa(i)
	}
	close(c)
}

func main() {
	var c = make(chan string)

	go add(c)

	var arr []string

	for s := range c {
		arr = append(arr, s)
	}

	fmt.Println(arr)
}
