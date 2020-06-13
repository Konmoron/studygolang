package main

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("num is", num)
}

func testArrays() {
	a := [3]int{10}
	fmt.Println("a is", a)
	b := [...]string{"a", "b", "c"}
	c := b
	c[0] = "c"
	fmt.Println("b is", b, "c is", c)
	num := [5]int{1, 2, 3, 4, 5}
	fmt.Println("num is", num)
	changeLocal(num)
	fmt.Println("num is", num)
	fmt.Println("len of num is", len(num))
	d := [...]float64{4.4, 100, 5.9, 10.3333}
	for i := 0; i < len(d); i++ {
		fmt.Printf("%d the element is %.2f\n", i, d[i])
	}
	sum := float64(0)
	for _, v := range d {
		sum += v
	}
	fmt.Println("sum of d is", sum)
}

func substactOne(numbers []int) {
	for i := range numbers {
		numbers[i]++
	}
}

func testSlice() {
	a := []int{1}
	fmt.Println("a is", a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b is", b)
	c := b[:]
	e := b[:]
	for i := range c {
		c[i]++
	}
	fmt.Println("b is", b)
	for i := range e {
		e[i]++
	}
	fmt.Println("b is", b)
	fruitArray := [...]string{"apple", "orange", "grape", "mango", "water-melon", "pine-apple"}
	fruitSlice := fruitArray[1:3]
	fmt.Println("fruitSlice is", fruitSlice)
	fmt.Printf("length of slice is %d and capacity is %d\n", len(fruitSlice), cap(fruitSlice))
	fruitSlice = fruitSlice[:cap(fruitSlice)]
	fmt.Println("fruitSlice is", fruitSlice)

	d := make([]int, 1)
	fmt.Println("d is", d)

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars is", cars, "len of cars is", len(cars), "and capacity is", cap(cars))
	cars = append(cars, "BYD")
	fmt.Println("cars is", cars, "len of cars is", len(cars), "and capacity is", cap(cars))
	cars = append(cars, "BYD")
	fmt.Println("cars is", cars, "len of cars is", len(cars), "and capacity is", cap(cars))
	cars = append(cars, "BYD")
	fmt.Println("cars is", cars, "len of cars is", len(cars), "and capacity is", cap(cars))
	cars = append(cars, "BYD")
	fmt.Println("cars is", cars, "len of cars is", len(cars), "and capacity is", cap(cars))
	cars = append(cars, "BYD")
	fmt.Println("cars is", cars, "len of cars is", len(cars), "and capacity is", cap(cars))

	var names []string
	if names == nil {
		fmt.Println("names is nil")
		names = append(names, "tom", "taylor")
		fmt.Println("names =", names)
	}

	nos := []int{5, 6, 7}
	fmt.Println("nos is", nos)
	substactOne(nos)
	fmt.Println("nos is", nos)
}

func main() {
	//testArrays()
	testSlice()
}
