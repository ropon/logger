package logger

var Log Logger = NewConsoleLogger("debug")

type LogCfg struct {
	Level     string  `json:"level"`
	FilePath  string  `json:"file_path"`
	FileName  string  `json:"file_name"`
	MaxSize   int64   `json:"max_size"`
	SplitFlag bool    `json:"split_flag"`
	TimeDr    float64 `json:"time_dr"`
}

var defaultLogCfg = LogCfg{
	Level:    "debug",
	FilePath: "./log",
	FileName: "log.log",
}

func InitLog(logCfgS ...*LogCfg) (err error) {
	logCfg := &defaultLogCfg
	if len(logCfgS) == 1 {
		logCfg = logCfgS[0]
	}
	if logCfg.MaxSize == 0 {
		logCfg.MaxSize = 1024 * 1024 * 1024
	}
	if logCfg.SplitFlag && logCfg.TimeDr == 0 {
		logCfg.TimeDr = 2 * 60
	}
	Log, err = NewFileLogger(logCfg)
	return
}

func Close() {
	Log.Close()
}

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

func Fatal(format string, args ...interface{}) {
	Log.Fatal(format, args...)
}

func Panic(format string, args ...interface{}) {
	Log.Panic(format, args...)
}
