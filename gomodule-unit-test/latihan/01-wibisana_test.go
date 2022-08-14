package unit_test

import (
	"testing"
)

func Test_palindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "palindrome got false",
			args: args{
				s: "meong",
			},
			want: false,
		},
		{
			name: "palindrome got true",
			args: args{
				s: "radar",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindrome(tt.args.s); got != tt.want {
				t.Errorf("palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConsonant(t *testing.T) {
	type args struct {
		check string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Got vokal",
			args: args{
				check: "a",
			},
			want: "vokal",
		},
		{
			name: "Got konsonan",
			args: args{
				check: "k",
			},
			want: "konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Consonant(tt.args.check); got != tt.want {
				t.Errorf("Consonant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_palindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindrome("radar")
	}
}

func BenchmarkConsonant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Consonant("k")
	}
}
