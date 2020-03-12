package logger

import (
    "testing"
)

func BenchmarkTrace(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Info("abc")
        // fmt.Println("abc")
        // log.Println("abc")
    }
}

func TestInfo(t *testing.T) {
    EnableColor(true)
    Debug("abc_debug")
    Info("abc_info")
    Warn("abc_warn")
    Error("abc_err")
    Panic("abc_painc")
}

func TestInfoDepth(t *testing.T) {
    a1()
}

func a1()  {
    b1()
}

func b1()  {
    InfoDepth(4, "abc")
}