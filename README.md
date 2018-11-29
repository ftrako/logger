# 日志打印

基于log.Println()上的二次开发，显示内容更丰富。

注意：调用系统log.SetFlags()会影响logger日志输出，logger需要系统log的flag清零


## 使用说明

* 使用 go get -u github.com/ftrako/logger 下载或更新

* 目前只支持控制台模式

* 使用示例：

`logger.Info("hello world!") // 打印INFO级别的日志`

* 关闭字体颜色示例：

`logger.GetDefaultLogger().Color = false`

* 关闭函数名：

`logger.GetDefaultLogger().Func = false`

* 设置日志级别：

`logger.GetDefaultLogger().Level = logger.LevelTrace`