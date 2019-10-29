package main

import (
	"fmt"
	"sync"
)

type rect1 struct {
	length int
	width  int
}

func (r rect1) aera1(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length <= 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}

	if r.width <= 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func main() {
	var wg sync.WaitGroup
	r1 := rect1{-123, 234}
	r2 := rect1{324, -234}
	r3 := rect1{324,34}
	r4 := rect1{-34, -23}
	rects := []rect1{r1, r2, r3, r4}
	for _, v := range rects {
		wg.Add(1)
		go v.aera1(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
