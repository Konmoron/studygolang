package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	ch <- "steven"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
