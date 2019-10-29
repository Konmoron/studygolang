package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func printTime() {
	fmt.Println(time.Now())
}

func main() {
	printTime()
	c := cron.New()
	err := c.AddFunc("* * * * * *", func() {
		printTime()
	})
	if err != nil {
		fmt.Println(err)
	}
	c.Start()
	time.Sleep(60*time.Second)
}
