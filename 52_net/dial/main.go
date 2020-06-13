package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addresses := []string{"www.baidu.com:80", "www.baidu.com:10246", "cannot-dial-domain-bala.com:10246"}
	// conn, err := net.DialTimeout("tcp", "dc-self-prod.yonyoucloud-k8s.com:10246", time.Second)
	for _, address := range addresses {
		conn, err := net.DialTimeout("tcp", address, time.Second)
		if err != nil {
			fmt.Println("net.DialTimeout:", err)
			continue
		}
		defer conn.Close()
		fmt.Println(conn.RemoteAddr())
		fmt.Println(conn.LocalAddr())
		// utils.PrintToJson(conn.LocalAddr)
	}

}
