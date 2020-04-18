## logger

+ 支持配置日志级别
+ 支持自定义拆分方式

-------

### 使用方法

```go
func main() {
    logger.InitLog("Info", "api")
    defer logger.Close()
    logger.Info("这是一条信息日志")
}
```

----

调用`logger.Error() 会自动执行os.Exit(1)`
调用`logger.Fatal() 会自动执行os.Panic()`