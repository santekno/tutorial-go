package helper

import "testing"

func Test_palindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindrome(tt.args.input); got != tt.want {
				t.Errorf("palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_palindrome(b *testing.B) {
	type args struct {
		input string
	}
	benchmark := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, bm := range benchmark {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				palindrome(bm.args.input)
			}
		})
	}
}
