package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		kata string
	}
	tests := []struct {
		name             string
		args             args
		wantIsPalindrome bool
	}{
		{
			name             : "test",
			args            : args{
				"katak",
			},
			wantIsPalindrome : true,
		},
		{
			name             : "test",
			args            : args{
				"kata",
			},
			wantIsPalindrome : false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsPalindrome := isPalindrome(tt.args.kata); gotIsPalindrome != tt.wantIsPalindrome {
				t.Errorf("isPalindrome() = %v, want %v", gotIsPalindrome, tt.wantIsPalindrome)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome("katak")
	}
}
