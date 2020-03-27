package logger

import (
	"fmt"
	"path"
	"runtime"
)

//获取函数名 文件名 行
func getCallerInfo(skip int) (funcName, fileName string, line int, err error) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		err = fmt.Errorf("%s", "获取程序信息失败")
		return
	}
	funcName = path.Base(runtime.FuncForPC(pc).Name())
	fileName = path.Base(fileName)
	return
}
