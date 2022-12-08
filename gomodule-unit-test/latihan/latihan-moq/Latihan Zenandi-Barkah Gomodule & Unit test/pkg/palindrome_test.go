package pkg

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success expected palindrom true",
			args: args{
				word: "madam",
			},
			want: true,
		},
		{
			name: "success expected palindrom false",
			args: args{
				word: "zenandi",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.word); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	word := "madam"
	for i := 0; i < b.N; i++ {
		IsPalindrome(word)
	}
}
