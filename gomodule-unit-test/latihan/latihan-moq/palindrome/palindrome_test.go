package palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Palindrome true",
			args: args{
				inputString: "kodok",
			},
			want: true,
		},
		{
			name: "Palindrome false",
			args: args{
				inputString: "unitest",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.inputString); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	inputString := "kodok"
	for i := 0; i < b.N; i++ {
		IsPalindrome(inputString)
	}
}
