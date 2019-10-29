package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// https://stackoverflow.com/questions/26290485/golang-yaml-reading-with-map-of-maps
// https://my.oschina.net/zhongzhong5/blog/1601607


func main() {
	fmt.Println("read yaml to map")
	var monitorApps = make(map[string]map[string][]string)
	buffer, err := ioutil.ReadFile("./45_config/config.yaml")
	if err == nil {
		err = yaml.Unmarshal(buffer, &monitorApps)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("monitor_apps:", monitorApps)
			fmt.Println("User-message", monitorApps["User-message"])
			fmt.Println("User-message[must]", monitorApps["User-message"]["must"][0])
		}
	} else {
		fmt.Println(err)
	}
}
