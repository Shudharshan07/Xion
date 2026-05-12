package main

import (
	"io"
	"log"
	"testing"
)

func Benchmark(b *testing.B) {
	log.SetOutput(io.Discard)
	data := NewData()
	b.ResetTimer()
	for b.Loop() {
		data.UpdateData()
	}
}
