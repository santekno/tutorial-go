// Package strrev exports functions for reversing strings.
package strrev // import "4d63.com/strrev"

import (
	"unicode"
	"unicode/utf8"
)

// Reverse reverses the given string, maintaining the bytes within any multibyte unicode characters in their existing order so that the character is still rendered correctly.
func Reverse(s string) string {
	reversed := make([]byte, len(s))
	i := 0

	for len(s) > 0 {
		r, size := utf8.DecodeLastRuneInString(s)
		s = s[:len(s)-size]
		i += utf8.EncodeRune(reversed[i:], r)
	}

	return string(reversed)
}

// ReverseBytes reverses the given byte slice, maintaining the bytes within any multibyte unicode characters in their existing order so that the character is still rendered correctly.
func ReverseBytes(b []byte) []byte {
	reversed := make([]byte, len(b))
	i := 0

	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		b = b[:len(b)-size]
		i += utf8.EncodeRune(reversed[i:], r)
	}

	return reversed
}

// ReverseCombining reverses the given string, maintaining the bytes within any multibyte and combining unicode characters in their existing order so that the character is still rendered correctly.
func ReverseCombining(s string) string {
	reversed := make([]byte, len(s))
	marks := make([]rune, 0)
	i := 0

	for len(s) > 0 {
		r, size := utf8.DecodeLastRuneInString(s)
		s = s[:len(s)-size]
		if unicode.IsMark(r) {
			marks = append(marks, r)
		} else {
			i += utf8.EncodeRune(reversed[i:], r)
			for m := len(marks) - 1; m >= 0; m-- {
				mark := marks[m]
				i += utf8.EncodeRune(reversed[i:], mark)
			}
			marks = marks[:0]
		}
	}

	for m := len(marks) - 1; m >= 0; m-- {
		mark := marks[m]
		i += utf8.EncodeRune(reversed[i:], mark)
	}

	return string(reversed)
}

// ReverseCombiningBytes reverses the given byte slice, maintaining the bytes within any multibyte and combining unicode characters in their existing order so that the character is still rendered correctly.
func ReverseCombiningBytes(b []byte) []byte {
	reversed := make([]byte, len(b))
	marks := make([]rune, 0)
	i := 0

	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		b = b[:len(b)-size]
		if unicode.IsMark(r) {
			marks = append(marks, r)
		} else {
			i += utf8.EncodeRune(reversed[i:], r)
			for m := len(marks) - 1; m >= 0; m-- {
				mark := marks[m]
				i += utf8.EncodeRune(reversed[i:], mark)
			}
			marks = marks[:0]
		}
	}

	for m := len(marks) - 1; m >= 0; m-- {
		mark := marks[m]
		i += utf8.EncodeRune(reversed[i:], mark)
	}

	return reversed
}
