package main

import "fmt"

func main() {
	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}
	if finger > 1 {
		fmt.Println("finger > 1")
	} else if finger > 2 {
		fmt.Println("finger > 2")
	}
}
