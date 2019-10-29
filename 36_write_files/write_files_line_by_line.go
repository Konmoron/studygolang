package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./36_write_files/test_line_by_line.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	d := []string{"Welcome to the world Go1.", "Go is a compiled language", "It is easy to learn Go"}
	for _, v := range d {
		_, err = fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file write successfully!")

	// write append
	f, err = os.OpenFile("./36_write_files/test_line_by_line.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy"
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")
}
