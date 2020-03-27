package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 日志结构体
type fileLogger struct {
	level    string
	filePath string
	fileName string
	file     *os.File
	errFile  *os.File
	maxSize  int64
	lastTime time.Time
}

//LogMsg 定义日志信息结构体
type logMsg struct {
	file    *os.File
	errFile *os.File
	level   string
	msg     string
}

//定义日志通道
var logChan = make(chan *logMsg, 500)

// NewFileLogger 日志结构体 构造函数
func NewFileLogger(level, filePath, fileName string) *fileLogger {
	fileLogger := &fileLogger{
		level:    level,
		filePath: filePath,
		fileName: fileName + ".log",
		maxSize:  10 * 1024 * 1024,
		lastTime: time.Now(),
	}
	_ = fileLogger.initFile()
	return fileLogger
}

//新建日志文件
func (f *fileLogger) initFile() error {
	logName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("打开日志文件%s异常, 报错:%v", logName, err)
	}
	f.file = fileObj
	errLogName := fmt.Sprintf("err_%s", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("打开日志文件%s异常, 报错:%v", errLogName, err)
	}
	f.errFile = errFileObj
	return nil
}

//再次封装写日志函数
func (f *fileLogger) wLog(level string, format string, args ...interface{}) {
	if getLevel(f.level) > getLevel(level) {
		return
	}
	// 按大小拆分日志
	// f.file = f.checkSplitLog(f.file)
	//f.checkSplitLog()
	msgInfo := fmt.Sprintf(format, args...)
	nowStr := time.Now().Local().Format("2006-01-02 15:04:05.000")
	funcName, fileName, line, _ := getCallerInfo(4)
	msg := fmt.Sprintf("[%s] [%s] [%s:%s] %d %s", nowStr, level, fileName, funcName, line, msgInfo)
	//将日志信息发送通道
	logMsgTemp := &logMsg{
		file:    f.file,
		errFile: f.errFile,
		level:   level,
		msg:     msg,
	}
	logChan <- logMsgTemp
	//_, _ = fmt.Fprintln(f.file, logMsg)
	//if Getlevel(logLevel) >= Getlevel("ERROR") {
	//	// 按大小拆分时需要
	//	// f.errFile = f.checkSplitLog(f.errFile)
	//	//_, _ = fmt.Fprintln(f.errFile, logMsg)
	//}
}

//将日志写入文件
func FileLog() {
	for logMsg := range logChan {
		//将日志写入文件
		_, _ = fmt.Fprintln(logMsg.file, logMsg.msg)
		if getLevel(logMsg.level) >= getLevel("ERROR") {
			_, _ = fmt.Fprintln(logMsg.errFile, logMsg.msg)
			switch getLevel(logMsg.level) {
			case getLevel("ERROR"):
				os.Exit(1)
			case getLevel("FATAL"):
				panic(logMsg.msg)
			}
		}
	}
}

//日志拆分
func (f *fileLogger) checkSplitLog() {
	// 按大小拆分
	// fileInfo, _ := file.Stat()
	// fileSize := fileInfo.Size()
	// if fileSize < f.maxSize {
	// 	return file
	// }
	// 按时间拆分 2小时拆分1次
	timeD := time.Now().Sub(f.lastTime).Hours()
	if timeD >= 2 {
		fileName := f.file.Name()
		backupName := fmt.Sprintf("%s_bak%v", fileName, time.Now().Unix())
		_ = f.file.Close()
		_ = os.Rename(fileName, backupName)
		fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Errorf("打开日志文件%s异常, 报错:%v", fileName, err))
		}

		errFileName := f.errFile.Name()
		errBackupName := fmt.Sprintf("%s_bak%v", errFileName, time.Now().Unix())
		_ = f.errFile.Close()
		_ = os.Rename(errFileName, errBackupName)
		errFileObj, err := os.OpenFile(errFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Errorf("打开日志文件%s异常, 报错:%v", errFileName, err))
		}
		// 将当前时间复制给上次拆分时间
		f.lastTime = time.Now()
		f.file = fileObj
		f.errFile = errFileObj
	}
}

// Debug 调试日志
func (f *fileLogger) Debug(format string, args ...interface{}) {
	f.wLog("DEBUG", format, args...)
}

// Info 一般日志
func (f *fileLogger) Info(format string, args ...interface{}) {
	f.wLog("INFO", format, args...)
}

// Warn 警告日志
func (f *fileLogger) Warn(format string, args ...interface{}) {
	f.wLog("WARN", format, args...)
}

// Error 错误日志
func (f *fileLogger) Error(format string, args ...interface{}) {
	f.wLog("ERROR", format, args...)
}

// Fatal 严重错误日志
func (f *fileLogger) Fatal(format string, args ...interface{}) {
	f.wLog("FATAL", format, args...)
}

// Close 关闭文件句柄
func (f *fileLogger) Close() {
	time.Sleep(time.Millisecond * 500)
	_ = f.file.Close()
	_ = f.errFile.Close()
}
