package logger

import "testing"

func TestLog(t *testing.T) {
	InitLog()
	defer Close()
	for {
		Log.Debug("这是测试Debug日志")
		Log.Info("这是测试Info日志")
		Log.Warn("这是测试Warn日志")
		Log.Error("这是测试Error日志")
		Log.Fatal("这是测试Fatal日志")
	}
}
