package main

import (
	"testing"
)

func Test_checkPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "should be palindrome",
			args: args{
				str: "madam",
			},
			want: true,
		},
		{
			name: "should not be palindrome",
			args: args{
				str: "syauqi",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPalindrome(tt.args.str); got != tt.want {
				t.Errorf("checkPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Should be Success Run Main Function",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
