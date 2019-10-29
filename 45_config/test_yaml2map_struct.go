package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// https://stackoverflow.com/questions/26290485/golang-yaml-reading-with-map-of-maps
// https://my.oschina.net/zhongzhong5/blog/1601607

type MonitorCount struct {
	MonitorString string `yaml:"monitor_string"`
	MaxNum        int    `yaml:"max_num"`
	CurrNum       int    `yaml:"curr_num"`
}

type MonitorApp1 struct {
	AppName  string         `yaml:"app_name"`
	LogFiles []string       `yaml:"log_files"`
	PreMatch string         `yaml:"pre_match"`
	Must     []string       `yaml:"must"`
	MustNot  []string       `yaml:"must_not"`
	Counts   []MonitorCount `yaml:"counts"`
}

type Config struct {
	Label       string                 `yaml:"label"`
	RunModel    string                 `yaml:"run_model"`
	IsDebug     bool                   `yaml:"debug"`
	MonitorApps map[string]MonitorApp1 `yaml:"monitor_apps"`
}

func main() {
	fmt.Println("read yaml to map")
	//var config = make(map[string]interface{})
	var config Config
	//var monitorApps = make(map[string]interface{})
	buffer, err := ioutil.ReadFile("./45_config/config2map_struct.yaml")
	if err == nil {
		err = yaml.Unmarshal(buffer, &config)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("monitor_apps:", config)
			fmt.Println("label:", config.Label)
			fmt.Println("run model:", config.RunModel)
			fmt.Println("is debug:", config.IsDebug)
			fmt.Println()
			if config.IsDebug {
				//fmt.Println("monitor_apps", config.MonitorApps)
				//fmt.Println("monitor app Rest:", config.MonitorApps["Rest"])
				//fmt.Println("monitor app Rest must not:", config.MonitorApps["Rest"].MustNot)
				//fmt.Println("monitor app Rest must not 0:", config.MonitorApps["Rest"].MustNot[0])
				//fmt.Println("count:", config.MonitorApps["Rest"].Counts[0].MaxNum * 4)
				//if len(config.MonitorApps["User-message"].Counts) == 0 {
				//	fmt.Println("User-message count: 0")
				//} else {
				//	fmt.Println("User-message count string:", config.MonitorApps["User-message"].Counts[0].MonitorString)
				//	fmt.Println("User-message count string:", config.MonitorApps["User-message"].Counts[0].MaxNum * 4)
				//}
				for appName := range config.MonitorApps {
					fmt.Println("app name:", appName)
					fmt.Println("log files:", config.MonitorApps[appName].LogFiles)
					fmt.Println("pre_match:", config.MonitorApps[appName].PreMatch)
					fmt.Println("must:", config.MonitorApps[appName].Must)
					fmt.Println("must not:", config.MonitorApps[appName].MustNot)
					//var counts []MonitorCount
					//counts = config.MonitorApps[appName].Counts
					//isContain := false
					//for _, c := range config.MonitorApps[appName].Counts {
					//	if c.MonitorString == "__all__" {
					//		isContain = true
					//		break
					//	}
					//}
					//if ! isContain {
					//	c := MonitorCount{MonitorString: "__all__", MaxNum:2}
					//	config.MonitorApps[appName].Counts = append(config.MonitorApps[appName].Counts, c)
					//}
					//if config.MonitorApps[appName].Counts == nil {
					//	//config.MonitorApps[appName].Counts = *make([]MonitorCount)
					//	//config.MonitorApps[appName].Counts = &counts
					//	fmt.Println("User-message count: 0")
					//	c := MonitorCount{MonitorString: "__all__", MaxNum:2}
					//	//counts = append(counts, c)
					//	//config.MonitorApps[appName].Counts = &counts
					//} else {
					//	counts = config.MonitorApps[appName].Counts
					//	//fmt.Println(appName, "count string:", config.MonitorApps[appName].Counts[0].MonitorString)
					//	//fmt.Println(appName, "count string:", config.MonitorApps[appName].Counts[0].MaxNum * 4)
					//	//config.MonitorApps[appName].Counts.
					//	isContain := false
					//	for _, c := range counts {
					//		if c.MonitorString == "__all__" {
					//			isContain = true
					//			break
					//		}
					//	}
					//	if ! isContain {
					//		c := MonitorCount{MonitorString: "__all__", MaxNum:2}
					//		counts = append(counts, c)
					//	}
					//}
					//config.MonitorApps[appName].Counts = counts
					for _, c := range config.MonitorApps[appName].Counts {
						c.MaxNum = 10
						fmt.Println(appName, "count string:", c.MonitorString)
						fmt.Println(appName, "count max num:", c.MaxNum)
						fmt.Println(appName, "count curr num:", c.CurrNum)
					}

					//config.MonitorApps[appName].Counts = &counts

					fmt.Printf("\n\n")
				}
			}
		}
	} else {
		fmt.Println(err)
	}
}
