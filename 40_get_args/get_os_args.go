package main

import (
	"fmt"
	"os"
)

const usage = `
用法: 
	脚本名 参数1 参数2`

func main() {
	args := os.Args
	//fmt.Println(args)
	if len(args) == 1 {
		fmt.Println(usage)
	} else {
		for index, arg := range args {
			fmt.Println("arg index:", index, "arg is:", arg)
		}
	}
}
