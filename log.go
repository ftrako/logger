package logger

var defaultLogger *Logger

func Trace(f interface{}, v ...interface{}) {
    defaultLogger.Trace(f, v...)
}

func Debug(f interface{}, v ...interface{}) {
    defaultLogger.Debug(f, v...)
}

func Info(f interface{}, v ...interface{}) {
    defaultLogger.Info(f, v...)
}

func Warn(f interface{}, v ...interface{}) {
    defaultLogger.Warn(f, v...)
}

func Error(f interface{}, v ...interface{}) {
    defaultLogger.Error(f, v...)
}

func Panic(f interface{}, v ...interface{}) {
    defaultLogger.Panic(f, v...)
}

func GetDefaultLogger() *Logger {
    return defaultLogger
}

func init() {
    defaultLogger = NewLogger()
}
