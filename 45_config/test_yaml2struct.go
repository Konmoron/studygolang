package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// https://stackoverflow.com/questions/26290485/golang-yaml-reading-with-map-of-maps
// https://my.oschina.net/zhongzhong5/blog/1601607

type MonitorApp struct {
	AppName string `yaml:"app_name"`
	LogFiles []string `yaml:"log_files"`
	PreMatch string `yaml:"pre_match"`
	Must []string `yaml:"must"`
	MustNot []string `yaml:"must_not"`
}


func main() {
	fmt.Println("read yaml to map")
	var monitorApps []MonitorApp
	var monitorAppsMap = make(map[string]MonitorApp)

	buffer, err := ioutil.ReadFile("./45_config/config2struct.yaml")
	if err == nil {
		err = yaml.Unmarshal(buffer, &monitorApps)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("monitor_apps:", monitorApps)
			for _, monitorApp := range monitorApps {
				monitorAppsMap[monitorApp.AppName] = monitorApp
			}
			fmt.Println("app_name", monitorAppsMap["User-message"].AppName)
			fmt.Println("must", monitorAppsMap["User-message"].Must[0])
			//fmt.Println("User-message", monitorApps["User-message"])
			//fmt.Println("User-message[must]", monitorApps["User-message"]["must"][0])
		}
	} else {
		fmt.Println(err)
	}
}
