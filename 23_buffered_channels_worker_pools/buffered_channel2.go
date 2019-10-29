package main

import (
	"fmt"
	"time"
)

func write(chnl chan int) {
	for i := 0; i < 5; i++ {
		chnl <- i
		fmt.Println("successfully wrote", i, "to chnl")
	}
	close(chnl)
}

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}
