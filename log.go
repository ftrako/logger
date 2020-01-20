package logger

var defaultLogger *Logger

// func Trace(f interface{}, v ...interface{}) {
//     defaultLogger.Trace(f, v...)
// }

// 参考log.LstdFlags
func SetFlag(flag int) {
    defaultLogger.Flag = flag
}

func EnableColor(enable bool) {
    defaultLogger.Color = enable
}

func EnablePrefix(enable bool) {
    defaultLogger.Prefix = enable
}

func EnableFuncName(enable bool) {
    defaultLogger.Func = enable
}

func SetLevel(level int) {
    defaultLogger.Level = level
}

func Debug(f interface{}, v ...interface{}) {
    defaultLogger.Debug(f, v...)
}

// depth应至少3及以上
func DebugDepth(depth int, f interface{}, v ...interface{}) {
    defaultLogger.DebugDepth(depth, f, v...)
}

func Info(f interface{}, v ...interface{}) {
    defaultLogger.Info(f, v...)
}

func InfoDepth(depth int, f interface{}, v ...interface{}) {
    defaultLogger.InfoDepth(depth, f, v...)
}

func Warn(f interface{}, v ...interface{}) {
    defaultLogger.Warn(f, v...)
}

func WarnDepth(depth int, f interface{}, v ...interface{}) {
    defaultLogger.WarnDepth(depth, f, v...)
}

func Error(f interface{}, v ...interface{}) {
    defaultLogger.Error(f, v...)
}

func ErrorDepth(depth int, f interface{}, v ...interface{}) {
    defaultLogger.ErrorDepth(depth, f, v...)
}

func Panic(f interface{}, v ...interface{}) {
    defaultLogger.Panic(f, v...)
}

func PanicDepth(depth int, f interface{}, v ...interface{}) {
    defaultLogger.PanicDepth(depth, f, v...)
}

// Deprecated
func GetDefaultLogger() *Logger {
    return defaultLogger
}

func init() {
    defaultLogger = NewLogger()
}
