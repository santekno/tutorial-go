// You can edit this code!
// Click here and start typing.
package main

import (
	"testing"
)

func Test_isVokal(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "is vokal",
			args: args{
				val: "a",
			},
			want: true,
		},
		{
			name: "is kosonan",
			args: args{
				val: "c",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVokal(tt.args.val); got != tt.want {
				t.Errorf("isVokal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeString(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

		{
			name: "is palindrome",
			args: args{
				val: "assa",
			},
			want: true,
		},
		{
			name: "not palindrome",
			args: args{
				val: "not",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeString(tt.args.val); got != tt.want {
				t.Errorf("isPalindromeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeInteger(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "is palindrome",
			args: args{
				val: 1331,
			},
			want: true,
		},
		{
			name: "not palindrome",
			args: args{
				val: 12345,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeInteger(tt.args.val); got != tt.want {
				t.Errorf("isPalindromeInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
