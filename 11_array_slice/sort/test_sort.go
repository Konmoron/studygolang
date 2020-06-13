package main

import (
	"fmt"
	"sort"
)

func main() {
	files := []string{"Makefile", "Test.conf", "util.go", "Makefile", "misc.go", "main.go"}

	target := "Makefile"
	fmt.Println(files)
	sort.Strings(files)
	fmt.Println(files)
	i := sort.SearchStrings(files, target)
	//i := sort.Search(len(files),
	//	func(i int) bool { return files[i] >= target })
	if i < len(files) && files[i] == target {
		fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
	} else {
		fmt.Println(i)
	}

	for i := range files {
		if files[i] >= target {
			fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
		}
	}
}