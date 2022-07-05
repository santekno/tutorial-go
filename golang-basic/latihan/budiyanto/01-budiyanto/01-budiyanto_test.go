package main

import (
	"testing"
)

func TestCheckPolindrome(t *testing.T) {
	cases := map[string]bool{
		"world": false,
		"madam": true,
		"madm":  false,
		"maam":  true,
		"mm":    true,
		"ms":    false,
		"s":     true,
	}

	for key, value := range cases {
		result := checkPolindrome(key)
		if result != value {
			t.Fatal("Expected", value, "got", result)
		}
	}
}
