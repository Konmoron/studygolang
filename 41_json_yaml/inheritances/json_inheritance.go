package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	NameA string `json:"name_a"`
}

type B struct {
	NameB string `json:"name_b"`
	A
}

type BSetNameForA struct {
	NameB string `json:"name_b"`
	TempA A
}

func main() {
	a := A{
		NameA: "a",
	}
	b1 := B{
		NameB: "b",
		A: a,
	}

	by, err := json.Marshal(&b1)
	if err == nil {
		fmt.Println(string(by))
	}
	b2 := BSetNameForA{
		NameB: "b",
		TempA: a,
	}
	by2, err := json.Marshal(&b2)
	if err == nil {
		fmt.Println(string(by2))
	}
}

