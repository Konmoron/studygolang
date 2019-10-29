package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x += 1
	m.Unlock()
	wg.Done()
}

func main() {
	startTime := time.Now()
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 10000; i++ {
		w.Add(1)
		go increment(&w, &m)
		//go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x is", x)
	endTime := time.Now()
	spendTime := endTime.Sub(startTime)
	fmt.Println("spend time is", spendTime)
	//time.Sleep(10*time.Second)
}
