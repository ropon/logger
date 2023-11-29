package logger

import "testing"

func init() {
	_ = InitLog()
}

func BenchmarkLogInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info("this is test log")
	}
}

func BenchmarkLogError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Error("this is test err log")
	}
}

func BenchmarkLogFatal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Error("this is test fatal log")
	}
}

func BenchmarkLogPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Error("this is test panic log")
	}
}
