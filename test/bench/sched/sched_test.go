package sched

import (
	"runtime"
	"strconv"
	"testing"
)

func BindM() []string {
	var s []string
	runtime.LockOSThread()
	for i := 0; i < 100000; i++ {
		s = append(s, strconv.Itoa(i))
	}
	return s
}

func BenchmarkSched(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 10; i++ {
		go BindM()
	}
}
