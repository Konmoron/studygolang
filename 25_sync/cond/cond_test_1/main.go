package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedRsc = false

func main() {
	//var wg sync.WaitGroup
	//wg.Add(2)
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		defer c.L.Unlock()
		for sharedRsc == false {
			fmt.Println("goroutine1 wait")
			c.Wait()
			//wg.Done()
			fmt.Println("goroutine1", sharedRsc)
		}
		time.Sleep(2*time.Second)
		fmt.Println("goroutine1", sharedRsc)
	}()

	go func() {
		// this go routine wait for changes to the sharedRsc
		c.L.Lock()
		defer c.L.Unlock()
		for sharedRsc == false {
			fmt.Println("goroutine2 wait")
			c.Wait()
			fmt.Println("goroutine2", sharedRsc)
			//wg.Done()
		}
		time.Sleep(1*time.Second)
		fmt.Println("goroutine2", sharedRsc)
	}()

	// this one writes changes to sharedRsc
	time.Sleep(2 * time.Second)
	c.L.Lock()
	fmt.Println("main goroutine ready")
	sharedRsc = true
	c.Broadcast()
	fmt.Println("main goroutine broadcast")
	c.L.Unlock()
	//wg.Wait()
	time.Sleep(5 * time.Second)
}
