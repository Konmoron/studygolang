package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) aera(wg *sync.WaitGroup) {
	if r.length <= 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		wg.Done()
		return
	}

	if r.width <= 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-123, 234}
	r2 := rect{324, -234}
	r3 := rect{324,34}
	r4 := rect{-34, -23}
	rects := []rect{r1, r2, r3, r4}
	for _, v := range rects {
		wg.Add(1)
		go v.aera(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
