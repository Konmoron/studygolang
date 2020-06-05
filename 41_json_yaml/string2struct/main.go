package main

import (
	"encoding/json"
	"fmt"
)

// 验证 输入的数据 比 结构体 的字段多
// json.Unmarshal 是否能够解析
//
// 验证结果：
// 输入数据 的字段比 结构体 的字段多，
// 多余的会被略过

type s struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	b := []byte(`{"name":"tom","age":10,"address":"Beijing"}`)

	var sObject s
	if err := json.Unmarshal(b, &sObject); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sObject)
}
