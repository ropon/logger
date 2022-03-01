package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

var (
	maxChanCount int64 = 50000
)

// FileLogger 日志结构体
type fileLogger struct {
	level     string
	filePath  string
	fileName  string
	file      *os.File
	errFile   *os.File
	maxSize   int64
	lastTime  time.Time
	splitFlag bool
	timeDr    float64
}

//LogMsg 定义日志信息结构体
type logMsg struct {
	file    *os.File
	errFile *os.File
	level   string
	msg     string
}

//定义日志通道
var logChan = make(chan *logMsg, maxChanCount)

// NewFileLogger 日志结构体 构造函数
func NewFileLogger(logCfg *LogCfg) (*fileLogger, error) {
	fileLogger := &fileLogger{
		level:     logCfg.Level,
		filePath:  logCfg.FilePath,
		fileName:  logCfg.FileName,
		maxSize:   logCfg.MaxSize,
		lastTime:  time.Now(),
		splitFlag: logCfg.SplitFlag,
		timeDr:    logCfg.TimeDr,
	}
	//判断日志目录是否存在
	_, err := os.Stat(logCfg.FilePath)
	isExistFlag := err == nil || os.IsExist(err)
	if !isExistFlag {
		err := os.MkdirAll(logCfg.FilePath, 0744)
		if err != nil {
			return fileLogger, err
		}
	}
	err = fileLogger.initFile()
	//启动协程将日志写入文件中
	go fileLogger.FileLog()
	return fileLogger, err
}

//新建日志文件
func (f *fileLogger) initFile() error {
	logName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("打开日志文件%s异常, 报错:%v", logName, err)
	}
	f.file = fileObj
	errLogName := fmt.Sprintf("%s.wf", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("打开错误日志文件%s异常, 报错:%v", errLogName, err)
	}
	f.errFile = errFileObj
	return nil
}

//再次封装写日志函数
func (f *fileLogger) wLog(level string, format string, args ...interface{}) {
	if getLevel(f.level) > getLevel(level) {
		return
	}
	msgInfo := fmt.Sprintf(format, args...)
	nowStr := time.Now().Local().Format("2006-01-02 15:04:05.000")
	funcName, fileName, line, _ := getCallerInfo(4)
	colorStr := colorMap[strings.ToLower(level)]
	baseInfo := fmt.Sprintf("%s:%s [grn:%d] %s:%d [%s]", level, nowStr, runtime.NumGoroutine(), fileName, line, funcName)
	msg := fmt.Sprintf(colorStr+" %s", baseInfo, msgInfo)
	//将日志信息发送通道
	logMsgTemp := &logMsg{
		file:    f.file,
		errFile: f.errFile,
		level:   level,
		msg:     msg,
	}
	select {
	case logChan <- logMsgTemp:
	default:

	}
}

// FileLog 将日志写入文件
func (f *fileLogger) FileLog() {
	defer f.file.Sync()
	defer f.errFile.Sync()
	for {
		//检查拆分日志
		f.checkSplitLog()
		select {
		case logMsg := <-logChan:
			//将日志写入文件
			_, _ = fmt.Fprintln(logMsg.file, logMsg.msg)
			if getLevel(logMsg.level) >= getLevel("ERRO") {
				_, _ = fmt.Fprintln(logMsg.errFile, logMsg.msg)
				switch getLevel(logMsg.level) {
				case getLevel("FATAL"):
					os.Exit(1)
				case getLevel("PANIC"):
					panic(logMsg.msg)
				}
			}
		}
	}
}

func reCrFile(file *os.File) *os.File {
	fileName := file.Name()
	backupName := fmt.Sprintf("%s_%v", fileName, time.Now().Local().Format("2006-01-02 15:04:05.000"))
	_ = file.Close()
	_ = os.Rename(fileName, backupName)
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s异常, 报错:%v", fileName, err))
	}
	return fileObj
}

//日志拆分
func (f *fileLogger) checkSplitLog() {
	if f.splitFlag {
		// 按时间拆分
		timeD := time.Now().Sub(f.lastTime).Minutes()
		if timeD >= f.timeDr {
			f.file = reCrFile(f.file)
			f.errFile = reCrFile(f.errFile)
			f.lastTime = time.Now()
		}
	} else {
		// 按大小拆分
		fileInfo, _ := f.file.Stat()
		fileSize := fileInfo.Size()
		if fileSize >= f.maxSize {
			f.file = reCrFile(f.file)
		}
		errFileInfo, _ := f.errFile.Stat()
		errFileSize := errFileInfo.Size()
		if errFileSize >= f.maxSize {
			f.errFile = reCrFile(f.errFile)
		}
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
	f.wLog("ERRO", format, args...)
}

// Fatal 严重错误日志
func (f *fileLogger) Fatal(format string, args ...interface{}) {
	f.wLog("FATAL", format, args...)
}

func (f *fileLogger) Panic(format string, args ...interface{}) {
	f.wLog("PANIC", format, args...)
}

func (f *fileLogger) Print(args ...interface{}) {
	s := fmt.Sprint(args...)
	f.wLog("DEBUG", "%s", s)
}

// Close 关闭文件句柄
func (f *fileLogger) Close() {
	_ = f.file.Close()
	_ = f.errFile.Close()
}
