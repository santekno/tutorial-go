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
		{
			name: "menentukan polindrome",
			args: args{
				input: "kasurinirusak",
			},
			want: true,
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

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("kasurinirusak")
	}
}
