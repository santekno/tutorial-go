package main

import (
	"testing"
)

func TestVocalOrNot(t *testing.T) {
	result := VocalOrNot("s")
	expect := true

	if result != expect {
		t.Fatal("Output tidak sesuai")
	}
}

func BenchmarkVocalOrNot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VocalOrNot("E")
	}
}
