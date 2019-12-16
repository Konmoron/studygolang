package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	l := rate.NewLimiter(3.333, 20)
	ctx := context.Background()
	start := time.Now()
	// 要处理二十个事件
	for i := 0; i < 400; i++ {
		//l.Wait(ctx)
		l.WaitN(ctx, 1)
		doTime := time.Now().String()
		fmt.Println("do something", i, doTime)
		// dosomething
	}
	fmt.Println(time.Since(start)) // output: 7.501262697s （初始桶内5个和每秒2个token）
}
