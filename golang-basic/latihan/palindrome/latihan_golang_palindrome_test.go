package main

import "testing"

func Test_isPalindromeWord(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "must palindrome",
			args: args{
				name: string("/baba/"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeWord(tt.args.name); got != tt.want {
				t.Errorf("isPalindromeWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
