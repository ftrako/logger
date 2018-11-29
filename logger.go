package logger

import (
    "log"
    "strconv"
    "time"
    "runtime"
    "fmt"
    "strings"
)

// log.Println()接口二次封装，接口更丰富

type Logger struct {
    Color     bool `json:"color"`      // 是否开启字体颜色
    Prefix    bool `json:"prefix"`     // 是否开启前缀，比如[INFO]
    Flag      int  `json:"flag"`       // 属性，参考log.LstdFlags
    Level     int  `json:"level"`      // 日志级别，参考Level类型
    CallDepth int  `json:"call_depth"` // 调用函数深度，特意提供可支持外部修改
}

func NewLogger() *Logger {
    log.SetFlags(0) // 不带任何属性，不用原系统属性
    b := new(Logger)
    b.Color = true
    b.Prefix = true
    b.Level = LevelTrace
    b.Flag = log.Lmicroseconds | log.Lshortfile
    b.CallDepth = 3
    return b
}

func (p *Logger) Trace(f interface{}, v ...interface{}) {
    p.write(LevelTrace, f, v...)
}

func (p *Logger) Debug(f interface{}, v ...interface{}) {
    p.write(LevelDebug, f, v...)
}

func (p *Logger) Info(f interface{}, v ...interface{}) {
    p.write(LevelInfo, f, v...)
}

func (p *Logger) Warn(f interface{}, v ...interface{}) {
    p.write(LevelWarn, f, v...)
}

func (p *Logger) Error(f interface{}, v ...interface{}) {
    p.write(LevelError, f, v...)
}

func (p *Logger) Panic(f interface{}, v ...interface{}) {
    p.write(LevelFatal, f, v...)
    panic(p.format(f, v...))
}

func (p *Logger) write(level Level, f interface{}, v ...interface{}) {
    if int(level) > p.Level { // 级别限制
        return
    }
    now := time.Now()

    caller := ""
    if p.Flag&(log.Lshortfile|log.Llongfile) != 0 {
        var ok bool
        _, file, line, ok := runtime.Caller(p.CallDepth)
        if !ok {
            file = "???"
            line = 0
        } else {
            if p.Flag&log.Lshortfile != 0 {
                short := file
                for i := len(file) - 1; i > 0; i-- {
                    if file[i] == '/' {
                        short = file[i+1:]
                        break
                    }
                }
                file = short
            }
        }
        caller = "[" + file + ":" + strconv.Itoa(line) + "] "
    }

    t := ""
    if p.Flag&log.Lmicroseconds != 0 {
        t = now.Format("2006-01-02 15:04:05.000") + " "
    } else if p.Flag&log.Ltime != 0 {
        t = now.Format("2006-01-02 15:04:05") + " "
    } else if p.Flag&log.Ldate != 0 {
        t = now.Format("2006-01-02") + " "
    }
    prefix := ""
    if p.Prefix {
        prefix = "[" + StrLevelMap[level] + "] "
    }
    str := t + prefix + caller + p.format(f, v...)
    if p.Color {
        str = colors[level](str)
    }
    log.Println(str)
}

func (p *Logger) format(f interface{}, v ...interface{}) string {
    var msg string
    switch f.(type) {
    case string:
        msg = f.(string)
        if len(v) == 0 {
            return msg
        }
        if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
            // format string
        } else {
            // do not contain format char
            msg += strings.Repeat(" %v", len(v))
        }
    default:
        msg = fmt.Sprint(f)
        if len(v) == 0 {
            return msg
        }
        msg += strings.Repeat(" %v", len(v))
    }
    return fmt.Sprintf(msg, v...)
}
