## logger

+ 支持配置日志级别
+ 支持自定义拆分方式
    + 默认按大小拆分 1G
    + 还支持按时间拆分

-------

### 使用方法

```go
func main() {
    logger.InitLog()
    logger.Info("这是一条信息日志")
}
```