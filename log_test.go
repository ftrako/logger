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
