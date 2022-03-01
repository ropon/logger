package main

import (
	"fmt"
	"github.com/ropon/logger"
	"time"
)

func main() {
	//默认按大小拆分，1G debug ./log log.log
	//err := logger.InitLog()

	//支持自定义
	logCfg := &logger.LogCfg{
		FilePath:  "./log",
		FileName:  "ropon.log",
		Level:     "debug",
		MaxSize:   100 * 1024 * 1024,
		SplitFlag: true,
		TimeDr:    1,
	}
	err := logger.InitLog(logCfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logger.Close()
	for i := 0; i < 10000000; i++ {
		logger.Debug("这是一条debug日志，我的名字是：%s，年龄：%d", "Ropon", 18)
		logger.Info("这是一条Info日志")
		logger.Warn("这是一条warn日志")
		//logger.Error("这是一条error日志 %s", "问题这么严重 。。。。")
		//logger.Panic("这是一条panic日志")
		//logger.Fatal("这是一条fatal日志")
		time.Sleep(time.Second)
	}
	select {}
}
