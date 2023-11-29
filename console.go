package logger

import (
	"fmt"
	"os"
	"time"
)

// ConsoleLogger 终端结构体
type ConsoleLogger struct {
	Level string
}

// NewConsoleLogger 构造终端结构体函数
func NewConsoleLogger(Level string) *ConsoleLogger {
	ConsoleObj := &ConsoleLogger{
		Level: Level,
	}
	return ConsoleObj
}

//再次封装写日志函数
func (f *ConsoleLogger) wLog(level string, format string, args ...interface{}) {
	if getLevel(f.Level) > getLevel(level) {
		return
	}
	msgInfo := fmt.Sprintf(format, args...)
	nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	funcName, fileName, line, _ := getCallerInfo(4)
	logMsg := fmt.Sprintf("%s:%s %s:%d [%s] %s", nowStr, level, fileName, line, funcName, msgInfo)
	_, _ = fmt.Fprintln(os.Stdout, logMsg)
}

// Debug 调试日志
func (f *ConsoleLogger) Debug(format string, args ...interface{}) {
	f.wLog("DEBUG", format, args...)
}

// Info 一般日志
func (f *ConsoleLogger) Info(format string, args ...interface{}) {
	f.wLog("INFO", format, args...)
}

// Warn 警告日志
func (f *ConsoleLogger) Warn(format string, args ...interface{}) {
	f.wLog("WARN", format, args...)
}

// Error 错误日志
func (f *ConsoleLogger) Error(format string, args ...interface{}) {
	f.wLog("ERROR", format, args...)
}

// Fatal 严重错误日志
func (f *ConsoleLogger) Fatal(format string, args ...interface{}) {
	f.wLog("FATAL", format, args...)
}

func (f *ConsoleLogger) Panic(format string, args ...interface{}) {
	f.wLog("PANIC", format, args...)
}

func (f *ConsoleLogger) Print(args ...interface{}) {
	s := fmt.Sprint(args...)
	f.wLog("DEBUG", "%s", s)
}

// Close 终端不需要关闭
func (f *ConsoleLogger) Close() {

}
