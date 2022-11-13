package main

import "testing"

func Test_palindrome(t *testing.T) {
	type args struct {
		inp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "palindrome",
			args: args{"madam"},
			want: true,
		},
		{
			name: "nog palindrome",
			args: args{"madarame"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindrome(tt.args.inp); got != tt.want {
				t.Errorf("palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
