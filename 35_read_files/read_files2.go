package main

import (
	"flag"
	"fmt"
)

func main() {
	fptr := flag.String("fpath", "test.txt", "file path read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)
}
