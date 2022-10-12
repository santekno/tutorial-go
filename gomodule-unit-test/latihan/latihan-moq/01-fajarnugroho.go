package main

import (
	"math/rand"
	"strings"
	"testing"
)

func IsPolindrom(param string) bool {
	var output string
	var matched bool

	matched = false
	for i := len(param) - 1; i >= 0; i-- {
		output += string(param[i])
	}

	if output == param {
		matched = true
	}
	return matched
}

func CheckIsVocal(param string) (string, int) {
	var isVocal string
	var status int

	if len(param) > 1 {
		status = 0
	} else {
		status = 1
	}

	switch strings.ToUpper(param) {
	case "A", "I", "U", "E", "O":
		isVocal = "vocal"
	default:
		isVocal = "konsonan"
	}
	return isVocal, status
}

func BenchmarkIsPolindrom(b *testing.B) {
	words := []string{"fajar", "nugroho", "ramukumar", "madam", "kayak"}
	for i := 0; i < b.N; i++ {
		IsPolindrom(words[rand.Intn(4)])
	}
}

func BenchmarkCheckIsVocal(b *testing.B) {
	x := []string{"a", "b", "e", "f", "g", "o"}
	for i := 0; i < b.N; i++ {
		CheckIsVocal(x[rand.Intn(5)])
	}
}
