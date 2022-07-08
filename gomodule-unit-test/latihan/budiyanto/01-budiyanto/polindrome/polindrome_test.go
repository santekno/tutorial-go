package main

import (
	"testing"
)

func Test_checkPolindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case mengeluarkan false",
			args: args{
				input: "world",
			},
			want: false,
		},
		{
			name: "test case mengeluarkan true",
			args: args{
				input: "madam",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPolindrome(tt.args.input); got != tt.want {
				t.Errorf("checkPolindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_checkPolindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPolindrome("world")
	}
}
