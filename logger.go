package logger

import "strings"

//日志库

//logLevelMap 定义日志级别map
var logLevelMap = map[string]int{
	"debug": 0,
	"info":  1,
	"warn":  2,
	"error": 3,
	"fatal": 4,
}

//Logger 是一个日志接口
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

// getLevel 获取日志级别
func getLevel(level string) int {
	//统一转成小写
	level = strings.ToLower(level)
	//map取不到返回对应类型零值
	return logLevelMap[level]
}

