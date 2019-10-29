package main

import (
	"fmt"
	"reflect"
)

type order1 struct {
	ordId      int
	customerId int
}

func createQuery1(q interface{}) {
	t := reflect.TypeOf(q)
	k := t.Kind()
	v := reflect.ValueOf(q)
	fmt.Println("Type:", t)
	fmt.Println("Kind:", k)
	fmt.Println("Value", v)
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i ++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}
}

func main() {
	o := order1{1,3}
	createQuery1(o)
}
