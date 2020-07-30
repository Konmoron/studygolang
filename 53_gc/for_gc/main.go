package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

var i int64 = 0

func main() {
	go func() {
		for {
			fmt.Println(atomic.LoadInt64(&i))
			time.Sleep(time.Millisecond * 500)
			runtime.GC() //手动触发GC
		}
	}()
	for {
		atomic.AddInt64(&i, 1)
	}
}
