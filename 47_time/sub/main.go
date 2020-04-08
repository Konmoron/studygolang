package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	time.Sleep(1000000000)
	t2 := time.Now()
	d := t2.Sub(t1)
	fmt.Printf("%d", d)
}
