package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		kat string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sukses Palindrome",
			args: args{
				kat :"madam",
			},
			want: bool(true),
		},
		{
			name: "gagal Palindrome",
			args: args{
				kat: "mada",
			},
			want: bool(false),
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.kat); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
type args struct {
		kat string
	}
	benchmarks := []struct {
		name string
		args args
	}{
		{
			name: "Palindrome (true)",
			args: args{kat: "madam"},
		},
		{
			name: "Palindrome (false)",
			args: args{kat: "rendy"},
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPalindrome(bm.args.kat)
			}
		})
	}
}
