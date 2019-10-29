package main

import "fmt"

type rectAreaError struct {
	err    string
	length int
	width  int
}

func (r rectAreaError) Error() string {
	return r.err
}

func (r rectAreaError) lengthNegative() bool {
	return r.length < 0
}

func (r rectAreaError) widthNegative() bool {
	return r.width < 0
}

func rectArea(length int, width int) (int, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err += "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &rectAreaError{err, length, width}
	}
	return length * width, nil
}

func main() {
	length, width := -20, -15
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*rectAreaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %d is less than zero\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: width %d is less than zero\n", err.width)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(area)
}
