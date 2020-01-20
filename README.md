# 日志打印

基于fmt.Println()上的二次开发，显示内容更丰富。

## 使用说明

* 使用 go get -u github.com/ftrako/logger 下载或更新

* 目前只支持控制台模式

* 使用示例：

`logger.Info("hello world!") // 打印INFO级别的日志`

* 关闭字体颜色示例：

`logger.EnableColor(false)`

* 关闭函数名：

`logger.EnableFuncName(false)`

* 设置日志级别：

`logger.SetLevel(logger.LevelInfo)`