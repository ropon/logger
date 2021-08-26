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
		FilePath:  "./log_dir",
		FileName:  "ropon.log",
		Level:     "debug",
		MaxSize:   100 * 1024 * 1024,
		SplitFlag: true,
		TimeDr:    2,
	}
	err := logger.InitLog(logCfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 1000000; i++ {
		logger.Debug("这是一条debug日志，我的名字是：%s，年龄：%d", "Ropon", 18)
		logger.Info("这是一条Info日志")
		logger.Warn("这是一条warn日志")
		logger.Error("这是一条error日志")
		time.Sleep(time.Second * 15)
	}
	select {}
}

/*
[2021-08-26 21:18:10.224] [DEBUG] [main.go:main.main] 24 [grn:2] 这是一条debug日志，我的名字是：Ropon，年龄：18
[2021-08-26 21:18:10.224] [INFO] [main.go:main.main] 25 [grn:2] 这是一条Info日志
[2021-08-26 21:18:10.224] [WARN] [main.go:main.main] 26 [grn:2] 这是一条warn日志
[2021-08-26 21:18:10.224] [ERROR] [main.go:main.main] 27 [grn:2] 这是一条warn日志
*/
