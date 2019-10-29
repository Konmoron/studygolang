package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("read file")
	f, err := os.Open("35_read_files/test.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	var offset int64 = 400
	var whence = 0
	newPosition, err := f.Seek(offset, whence)
	fmt.Println(newPosition)
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	//b := make([]byte, 400)
	//n, _ := f.ReadAt(b, newPosition)
	//fmt.Println(n)
	//fmt.Println(string(b[:n]))
	//fInfo, err := f.Stat()
	//if err != nil {
	//	fmt.Println(err)
	//}
	////fmt.Println(fInfo)
	//fmt.Println(fInfo.Size(), fInfo.Sys())

	// references:
	// 	1.https://segmentfault.com/a/1190000000377326
	//	2.https://ispycode.com/GO/Files-And-Directories/Seek-Positions-in-File
	//	3.https://stackoverflow.com/questions/38631982/golang-file-seek-and-file-writeat-not-working-as-expected
}
