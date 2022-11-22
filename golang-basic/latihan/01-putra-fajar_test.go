package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "must return true",
			args: args{
				input: "katak",
			},
			want: true,
		},
		{
			name: "must return false",
			args: args{
				input: "kucing",
			},
			want: false,
		},
		{
			name: "uppercase with palindrome word will return true",
			args: args{
				input: "Kodok",
			},
			want: true,
		},
		{
			name: "1 letter will return true",
			args: args{
				input: "a",
			},
			want: true,
		},
		{
			name: "input 2 letter will return false",
			args: args{
				input: "da",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.input); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPalindrome(b *testing.B) {
	input := "katak"
	for i := 0; i < b.N; i++ {
		isPalindrome(input)
	}
}
