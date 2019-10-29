package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timeout")
		} else if err.Temporary() {
			fmt.Println("Temporary error")
		} else {
			fmt.Println("generic error:", err)
		}
		return
	}
	fmt.Println(addr)
}
