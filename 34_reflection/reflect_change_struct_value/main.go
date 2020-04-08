package main

import (
	"fmt"
	"reflect"

	"imer.io/tools/common/utils"
)

// Foo struct
type Foo struct {
	Number int
	Text   string
}

// References:
// 	1. https://www.jianshu.com/p/40cb706853cd
// 	2. https://stackoverflow.com/questions/6395076/using-reflect-how-do-you-set-the-value-of-a-struct-field
func setValueByReflect(f *Foo) {
	fields := reflect.TypeOf(*f)
	values := reflect.ValueOf(f)
	// fmt.Println("value Kind after reflect.ValueOf(f) is ", values.Kind())

	if values.Kind() != reflect.Ptr || !values.Elem().CanSet() {
		fmt.Println("输入的对象是不可修改的")
		return
	} else {
		values = values.Elem()
		// fmt.Println("value Kind after values.Elem() is ", values.Kind())
	}

	for i := 0; i < values.NumField(); i++ {
		// field := fields.Field(i)
		fieldName := fields.Field(i).Name
		// fieldType := field.Type.Kind()
		index := []int{i}
		f := values.FieldByIndex(index)
		if !f.IsValid() {
			fmt.Println("没有找到Name的字段")
			return
		}

		// fmt.Println(f.Type().Name())

		fieldType := f.Type().Kind()

		switch fieldType {
		case reflect.String:
			fmt.Println("type of", fieldName, "is string, set to hello reflect")
			f.SetString("hello reflect")
		case reflect.Int:
			fmt.Println("type of", fieldName, "is int, set to 4444")
			f.SetInt(4444)
		default:
			fmt.Println("cannot match")
		}
	}
}

func main() {
	foo := Foo{123, "Hello world"}
	fmt.Println("origin foo is")
	utils.PrintToJson(foo)
	// reflect.ValueOf(&foo).Elem().Field(0).SetInt(321)
	// reflect.ValueOf(&foo).Elem().Field(1).SetString("hello reflect")

	setValueByReflect(&foo)

	fmt.Println("after reflect, foo is")
	utils.PrintToJson(foo)

	// fmt.Println(int(reflect.ValueOf(foo).Field(0).Int()))

	// reflect.ValueOf(&foo).Elem().Field(0).SetInt(321)

	// fmt.Println(int(reflect.ValueOf(foo).Field(0).Int()))
}
