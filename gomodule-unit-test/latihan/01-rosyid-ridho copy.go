package main

import (
	"fmt"
	"testing"
)

func BenchmarkIsPalindrom(b *testing.B) {
	palindrom := isPalindrom("katak")
	for i := 0; i < b.N; i++ {
		fmt.Printf("palindrom: %v\n", palindrom)
	}
}

func BenchmarkIsVocal(b *testing.B) {

	for i := 0; i < b.N; i++ {
		isVokal('a')
	}
}
