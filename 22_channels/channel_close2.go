package main

import "fmt"

func producer1(chnl chan int) {
	for i := 0; i <= 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go producer1(ch)
	for v := range ch {
		fmt.Println("Received ", v)
	}
}