package main

import "fmt"

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares(number int, squareOp chan int) {
	sum := 0
	//for number != 0 {
	//	digit := number % 10
	//	sum += digit * digit
	//	number /= 10
	//}
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareOp <- sum
}

func calcCubes(number int, cubeOp chan int) {
	sum := 0
	//for number != 0 {
	//	digit := number % 10
	//	sum += digit * digit * digit
	//	number /= 10
	//}
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeOp <- sum
}

func main() {
	number := 589
	sqrCh := make(chan int)
	cubeCh := make(chan int)
	go calcSquares(number, sqrCh)
	go calcCubes(number, cubeCh)
	squares, cubes := <-sqrCh, <-cubeCh
	fmt.Println("Final output is", squares+cubes)
}
