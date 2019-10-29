package main

import (
	"encoding/json"
	"fmt"
	"op-alarm/common/utils"
)

type A struct {
	NameA string `json:"name_a"`
}

type B struct {
	NameB string `json:"name_b"`
	A
}

func main() {
	a := A{
		NameA: "a",
	}
	b := B{
		NameB: "b",
		A: a,
	}

	utils.PrintToJson(b)

	by, err := json.Marshal(&b)
	if err == nil {
		fmt.Println(string(by))
	}
}
