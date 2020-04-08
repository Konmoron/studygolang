package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("Type:%T Value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("Type:%T Value:%v\n", y, y)
}
