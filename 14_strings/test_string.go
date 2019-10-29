package main

import (
	"fmt"
)

func printCharsAndBytes(s string) {
	for index, rune := range s {
		//fmt.Println("index ", index, "chars", rune)
		fmt.Printf("index %d char %c\n", index, rune)
	}
}

func mutate(s string) string {
	rune := []rune(s)
	rune[0] = 'a'
	return string(rune)
}

func main() {
	// 遍历字符串
	name := "Señor"
	printCharsAndBytes(name)

	h := "hello"
	fmt.Println(mutate(h))
}
