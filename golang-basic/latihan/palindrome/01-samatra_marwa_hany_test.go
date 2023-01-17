package main

import "testing"

func Test_palindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "vokal",
			args: args{
				input: "madam"
			},
			want: "palindrome",
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			palindrome(tt.args.input)
		})
	}
}
