package main

import "fmt"

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest...")
	max := nums[0]
	for _, v := range nums {
		if v >= max {
			max = v
		}
	}
	fmt.Println("Largest Num in", nums, "is", max)
	defer finished()
}

func main() {
	nums := []int{123, 234, 89, 234, 78324}
	largest(nums)
}