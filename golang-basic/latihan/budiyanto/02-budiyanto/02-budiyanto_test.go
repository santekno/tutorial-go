package main

import (
	"testing"
)

func TestCheckVowelOrConsonant(t *testing.T) {
	casesConsonant := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z", "B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z"}
	casesVowel := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	casesUnknown := []string{"1", "2", "3"}

	for _, value := range casesConsonant {
		result := checkVowelOrConsonant(value)
		if result != "Konsonan" {
			t.Fatal("Expected Konsonan", "got", result)
		}
	}

	for _, value := range casesVowel {
		result := checkVowelOrConsonant(value)
		if result != "Vokal" {
			t.Fatal("Expected Vokal", "got", result)
		}
	}

	for _, value := range casesUnknown {
		result := checkVowelOrConsonant(value)
		if result != "Unknown" {
			t.Fatal("Expected Unknown", "got", result)
		}
	}
}
