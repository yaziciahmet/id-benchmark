package main

import (
	"runtime"
	"testing"

	nanoid "github.com/matoous/go-nanoid/v2"
)

const (
	alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	idLength = 16
)

func genNID() string {
	id, _ := nanoid.New()
	return id
}

func BenchmarkGenNID(b *testing.B) {
	b.ResetTimer()
	b.Run("genNID", func(b *testing.B) {
		for i := 0; i < c; i++ {
			genNID()
		}

		runtime.GC()
	})
}
