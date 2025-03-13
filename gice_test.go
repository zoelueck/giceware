package main

import (
	"testing"
)

func BenchmarkGeneratePhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generatePhrase(100000, 100000, "-")
	}
}

func TestSimple(t *testing.T) {
	if 1 != 1 {
		t.Error("expected 1 to equal 1")
	}
}
