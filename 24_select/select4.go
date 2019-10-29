package main

import (
	"fmt"
	"time"
)

func server3(ch chan string) {
	ch <- "from server3"
}

func server4(ch chan string) {
	ch <- "from server4"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server3(output1)
	go server4(output2)
	time.Sleep(1 * time.Second)
	select {
	case v1 := <-output1:
		fmt.Println(v1)
	case v2 := <-output2:
		fmt.Println(v2)
	}
}
