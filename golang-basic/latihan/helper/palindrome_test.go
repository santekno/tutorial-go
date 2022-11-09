package helper

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "must be palindrome",
			args: args{input: "madam"},
			want: true,
		},
		{
			name: "must not be a palindrome",
			args: args{input: "rendy"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.input); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	type args struct {
		input string
	}
	benchmarks := []struct {
		name string
		args args
	}{
		{
			name: "Palindrome (true)",
			args: args{input: "madam"},
		},
		{
			name: "Palindrome (false)",
			args: args{input: "rendy"},
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsPalindrome(bm.args.input)
			}
		})
	}
}
