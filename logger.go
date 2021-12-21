package logger

import "strings"

//日志库

const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

var (
	//logLevelMap 定义日志级别map
	logLevelMap = map[string]int{
		"debug": 0,
		"info":  1,
		"warn":  2,
		"erro":  3,
		"fatal": 4,
		"panic": 5,
	}
	colorMap = map[string]string{
		"debug": Blue + "%s" + Reset,
		"info":  Green + "%s" + Reset,
		"warn":  Yellow + "%s" + Reset,
		"erro":  Red + "%s" + Reset,
		"fatal": RedBold + "%s" + Reset,
		"panic": YellowBold + "%s" + Reset,
	}
)

func ColorForStatus(code int) string {
	switch {
	case code >= 200 && code <= 299:
		return Green
	case code >= 300 && code <= 399:
		return White
	case code >= 400 && code <= 499:
		return Yellow
	default:
		return Red
	}
}

func ColorForMethod(method string) string {
	switch {
	case method == "GET":
		return Blue
	case method == "POST":
		return Cyan
	case method == "PUT":
		return Yellow
	case method == "DELETE":
		return Red
	case method == "PATCH":
		return Green
	case method == "HEAD":
		return Magenta
	case method == "OPTIONS":
		return White
	default:
		return Reset
	}
}

//Logger 是一个日志接口
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Panic(format string, args ...interface{})
	Print(args ...interface{})
	Close()
}

// getLevel 获取日志级别
func getLevel(level string) int {
	//统一转成小写
	level = strings.ToLower(level)
	//map取不到返回对应类型零值
	return logLevelMap[level]
}
