package main

import (
	"fmt"
	"sync"
	"time"
)

var x1 = 0

func increment2(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x1 += 1
	<- ch
	wg.Done()
}

func main() {
	startTime := time.Now()
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 10000; i++ {
		w.Add(1)
		go increment2(&w, ch)
	}
	w.Wait()
	fmt.Println("final value is", x1)
	endTime := time.Now()
	spendTime := endTime.Sub(startTime)
	fmt.Println("spend time", spendTime)
}
