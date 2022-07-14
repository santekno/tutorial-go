package palindrome

import "testing"

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
			name: "katak must true",
			args: args{
				word: "katak",
			},
			want: true,
		},
		{
			name: "itik must false",
			args: args{
				word: "itik",
			},
			want: false,
		},
		{
			name: "madam must true",
			args: args{
				word: "madam",
			},
			want: true,
		},
		{
			name: "madam must true",
			args: args{
				word: "madam",
			},
			want: true,
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
