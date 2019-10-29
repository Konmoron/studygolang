package main

import (
	"fmt"
	"time"
)

func main() {
	timeLayout := "2006-01-02T15:04:05.999999999+08:00"
	timeStr := "2019-10-29T09:26:26.101245303+08:00"
	t, err := time.ParseInLocation(timeLayout, timeStr, time.Local)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(timeStr, t)
}
