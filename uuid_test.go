package main

import (
	"runtime"
	"testing"

	"github.com/google/uuid"
)

var (
	testUUIDStr = "b943846e-dc8f-4726-bf8f-f7169ca264e7"
	testUUID    = uuid.MustParse(testUUIDStr)
)

func genUUID() uuid.UUID {
	return uuid.New()
}

func marshalUUID(id string) uuid.UUID {
	return uuid.MustParse(id)
}

func unmarshalUUID(id uuid.UUID) string {
	return id.String()
}

func BenchmarkGenUUID(b *testing.B) {
	b.ResetTimer()
	b.Run("genUUID", func(b *testing.B) {
		for i := 0; i < c; i++ {
			genUUID()
		}

		runtime.GC()
	})
}

func BenchmarkMarshalUUID(b *testing.B) {
	b.ResetTimer()
	b.Run("marshalUUID", func(b *testing.B) {
		for i := 0; i < c; i++ {
			marshalUUID(testUUIDStr)
		}

		runtime.GC()
	})
}

func BenchmarkUnmarshalUUID(b *testing.B) {
	b.ResetTimer()
	b.Run("unmarshalUUID", func(b *testing.B) {
		for i := 0; i < c; i++ {
			unmarshalUUID(testUUID)
		}

		runtime.GC()
	})
}
