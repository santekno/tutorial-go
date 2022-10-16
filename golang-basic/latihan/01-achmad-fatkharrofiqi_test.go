package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should be palindrome",
			args: args{
				word: "madam",
			},
			want: true,
		},
		{
			name: "should not be palindrome",
			args: args{
				word: "achmad",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.word); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
