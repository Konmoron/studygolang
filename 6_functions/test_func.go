package main

import "fmt"

func caculateBill(price, no int) int {
	totalPrice := price*no
	return totalPrice
}

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

func getParams(param1 string, param2 string) {
	params := `param1: $1
param2: $2`
	fmt.Println(params, param1, param2)
}

func main()  {
	price := 50
	no := 20
	totalPrice := caculateBill(price, no)
	fmt.Println("total price is", totalPrice)
	area, perimeter := rectProps(10.3, 8.7)
	fmt.Println("area is", area, "perimeter is", perimeter)

	getParams("string1", "string2")
}
