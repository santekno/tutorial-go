package main

import (
	"testing"
)

func TestCheckVowelOrConsonant(t *testing.T) {
	cases := map[string]string{
		"a": "Vokal",
		"e": "Vokal",
		"i": "Vokal",
		"o": "Vokal",
		"u": "Vokal",
		"b": "Konsonan",
		"c": "Konsonan",
		"d": "Konsonan",
		"f": "Konsonan",
		"g": "Konsonan",
		"h": "Konsonan",
		"j": "Konsonan",
		"k": "Konsonan",
		"l": "Konsonan",
		"m": "Konsonan",
		"n": "Konsonan",
		"p": "Konsonan",
		"q": "Konsonan",
		"r": "Konsonan",
		"s": "Konsonan",
		"t": "Konsonan",
		"v": "Konsonan",
		"w": "Konsonan",
		"x": "Konsonan",
		"y": "Konsonan",
		"z": "Konsonan",
		"1": "Unknown",
	}

	for key, value := range cases {
		result := checkVowelOrConsonant(key)
		if result != value {
			t.Fatal("Expected", value, "got", result)
		}
	}
}
