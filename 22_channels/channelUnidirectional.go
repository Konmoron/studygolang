package main

import "fmt"

func sendData(sendch chan<- int) {
	fmt.Println("sendData: before send data")
	sendch <- 10
	fmt.Println("sendData: after send data")
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println("main: before send data")
	fmt.Println(<-chnl)
	fmt.Println("main: after send data")
}