package test

import "testing"


func Test_palindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	} {
		{
			name : "palindrome ?",
			args : args{
				input : "madam"
			},
			want : true
		},
		{
			name : "palindrome ?",
			args : args{
				input : "abab"
			},
			want : false
		},
	}
	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsPalindrome := isPalindrome(tt.args.kata); gotIsPalindrome != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", gotIsPalindrome, tt.want)
			}
		}) 
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i <b.N; i++ {
		isPalinrome("madam")
	}
}