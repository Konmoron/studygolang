package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./35_read_files/test.txt")
	if err != nil {
		fmt.Println("File reading err", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
