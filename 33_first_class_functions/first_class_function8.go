package main

import "fmt"

func iMap(s []int, f func(i int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func main() {
	s := []int{1, 3, 5, 8}
	i := iMap(s, func(i int) int {
		return i * 5
	})

	fmt.Println(i)
}
