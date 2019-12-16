package main

import (
	"encoding/json"
	"fmt"
	"imer.io/tools/common/utils"
)

type A struct {
	NameA string `json:"name_a"`
}

type B struct {
	NameB string `json:"name_b"`
	*A
}

func main() {
	a := A{
		NameA: "a",
	}
	b := B{
		NameB: "b",
		A:     &a,
	}

	by, err := json.Marshal(&b)
	if err == nil {
		fmt.Println(string(by))
	}

	s := `{"name_a": "a", "name_b": "b"}`
	var b1 B
	if err := json.Unmarshal([]byte(s), &b1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b1.NameA)
		utils.PrintToJson(b1)
	}
}
