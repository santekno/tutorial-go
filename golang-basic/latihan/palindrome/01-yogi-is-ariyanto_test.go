package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		words string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "success validate word palindrome",
			args: args{"katak"},
			want: true,
		},
		{
			name: "success validate word palindrome",
			args: args{"bangau"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.words); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
