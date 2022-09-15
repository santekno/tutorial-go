package main

import "testing"

func Test_isPalindrom(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Word is Palindrome",
			args: args{
				val: "Madam",
			},
			want: true,
		},
		{
			name: "Word isn't Palindrome",
			args: args{
				val: "Tree",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrom(tt.args.val); got != tt.want {
				t.Errorf("isPalindrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPalindrom(t *testing.B) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Word is Palindrome",
			args: args{
				val: "Madam",
			},
			want: true,
		},
		{
			name: "Word isn't Palindrome",
			args: args{
				val: "Tree",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := isPalindrom(tt.args.val); got != tt.want {
				t.Errorf("isPalindrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
