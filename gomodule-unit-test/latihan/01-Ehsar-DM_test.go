package main

import (
	"testing"
)

func Test_polindrome(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is polindrome",
			args: args{
				word: "madam",
			},
			want: true,
		},
		{
			name: "is not polindrome",
			args: args{
				word: "majas",
			},
			want: false,
		},
		{
			name: "is polindrome",
			args: args{
				word: "apa",
			},
			want: true,
		},
		{
			name: "is not polindrome",
			args: args{
				word: "dakar",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := polindrome(tt.args.word); got != tt.want {
				t.Errorf("polindrome() = %v, want %v", got, tt.want)
			}

		})
	}
}
