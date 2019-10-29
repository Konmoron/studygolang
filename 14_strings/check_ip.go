package main

import (
	"fmt"
	"regexp"
	"strings"
)

func validIP4(ipAddress string) bool {
	// 0.0.0.0 ~ 255.255.255.255
	ipAddress = strings.Trim(ipAddress, " ")

	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ipAddress) {
		return true
	}
	return false
}

func main() {
	ip1 := "192.168.0.0"

	vIP1 := validIP4(ip1)

	fmt.Printf("%s is a valid IP4 address : %v \n", ip1, vIP1)

	//-----------------------------

	ip2 := "0.168.0.0"

	vIP2 := validIP4(ip2)

	fmt.Printf("%s is a valid IP4 address : %v \n", ip2, vIP2)

	//-----------------------------

	fullIP6 := "FE80:0000:0000:0000:0202:B3FF:FE1E:8329"

	vfullIP6 := validIP4(fullIP6)

	fmt.Printf("%s is a valid IP4 address : %v \n", fullIP6, vfullIP6)

	//-----------------------------

	collapsedIP6 := "FE80::0202:B3FF:FE1E:8329"

	vcollapsedIP6 := validIP4(collapsedIP6)

	fmt.Printf("%s is a valid IP4 address : %v \n", collapsedIP6, vcollapsedIP6)

	//-----------------------------

	urlIP6 := "http://[2001:db8:0:1]:80"

	vurlIP6 := validIP4(urlIP6)

	fmt.Printf("%s is a valid IP4 address : %v \n", urlIP6, vurlIP6)

}
