package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	//// 设置日志格式为json格式
	////log.SetFormatter(&log.JSONFormatter{})
	//
	//// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	//// 日志消息输出可以是任意的io.writer类型
	//log.SetOutput(os.Stdout)
	//
	//// 设置日志级别为warn以上
	//log.SetLevel(log.InfoLevel)
}

func main() {
	//log.Debug("Useful debugging information.")
	//log.Info("Something noteworthy happened!")
	//log.Warn("You should probably take a look at this.")
	//log.Error("Something failed but I'm not quitting.")
	////log.Fatal("Bye.")   //log之后会调用os.Exit(1)
	//log.Panic("I'm bailing.")   //log之后会panic()
	//log.WithFields(log.Fields{
	//	"animal": "walrus",
	//	"size":   10,
	//}).Info("A group of walrus emerges from the ocean")
	//
	//log.WithFields(log.Fields{
	//	"omg":    true,
	//	"number": 122,
	//}).Warn("The group's number increased tremendously!")
	//
	//log.WithFields(log.Fields{
	//	"omg":    true,
	//	"number": 100,
	//}).Fatal("The ice breaks!")

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
