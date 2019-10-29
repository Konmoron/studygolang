package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println("write file")
	f, err := os.Create("./36_write_files/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("hello world")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes write successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}