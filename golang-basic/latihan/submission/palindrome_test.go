package submission

import (
	"math/rand"
	"testing"
	"time"
)

func Test_IsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "expect return true",
			args: args{
				s: "madam",
			},
			want: true,
		},
		{
			name: "expect return false",
			args: args{
				s: "hikayat",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_Palindrome(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IsPalindrome(ramdomString(5))
	}
}

func ramdomString(length int) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
