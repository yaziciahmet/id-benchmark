package main

import (
	"runtime"
	"testing"

	"github.com/lucsky/cuid"
)

func genCUID() string {
	return cuid.New()
}

func BenchmarkGenCUID(b *testing.B) {
	b.ResetTimer()
	b.Run("genCUID", func(b *testing.B) {
		for i := 0; i < c; i++ {
			genCUID()
		}

		runtime.GC()
	})
}
