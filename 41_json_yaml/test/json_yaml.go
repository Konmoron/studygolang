package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
)

type TestJsonYaml struct {
	Id int `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
}

func main() {
	str1 := `id: 6
name: "tom"`

	var t1 TestJsonYaml
	err := yaml.Unmarshal([]byte(str1), &t1)
	if err == nil {
		fmt.Println("yaml format:")
		fmt.Println("\tid:", t1.Id)
		fmt.Println("\tname:", t1.Name)
	} else {
		fmt.Println(err)
	}
	
	str2 := `{"id": 5, "name": "jack"}`
	var t2 TestJsonYaml
	err = json.Unmarshal([]byte(str2), &t2)
	if err == nil {
		fmt.Println("json formant:")
		fmt.Println("\tid:", t2.Id)
		fmt.Println("\tname:", t2.Name)
	} else {
		fmt.Println(err)
	}
}
