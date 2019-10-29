package main

import "fmt"

type testNew struct {
	keyword string
}

func newTestNew() *testNew {
	return &testNew{
		keyword: "hello",
	}
}

func test() {
	t1 := newTestNew()
	fmt.Printf("pointer of t1 is %p\n", t1)
}

func main() {
	t1 := newTestNew()
	t2 := newTestNew()

	fmt.Printf("pointer of t1 is %p\n", t1)
	fmt.Printf("pointer of t2 is %p\n", t2)

	test()
}
