package logger

import (
	"fmt"
	"time"
)

var Log Logger

func InitLog() {
	Log = NewFileLogger("Info", "./", "niuPi")
	go FileLog()
}

func Close() {
	Log.Close()
}

func ErrorExit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		Log.Error(err.Error())
		time.Sleep(time.Millisecond * 500)
	}
}

//再次封装
func Debug(format string, args ...interface{}) {
	Log.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	Log.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	Log.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	Log.Error(format, args...)
}
