package main

import (
	"log"
	"reflect"
)

type TestRflection struct {
	firstName string
	age       int
}

type TestUseRflection struct {
	testRflection *TestRflection
}

func main() {
	//t1 := TestRflection{firstName: "tom", age: 10}
	t := TestUseRflection{}
	t1 := t.testRflection
	tType := reflect.TypeOf(t1)
	log.Println(tType)
}