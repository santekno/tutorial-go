package main

import (
	"math/rand"
	"testing"
)

func BenchmarkIsPolindrom(b *testing.B) {
	words := []string{"fajar", "nugroho", "ramukumar", "madam", "kayak"}
	for i := 0; i < b.N; i++ {
		IsPolindrom(words[rand.Intn(4)])
	}
}
