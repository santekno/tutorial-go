package latihan

import "testing"

func Test_checkPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		testName string
		args     args
		expect   bool
	}{
		{
			testName: "it should return false if given 'cosmic' input",
			args: args{
				str: "cosmic",
			},
			expect: false,
		},
		{
			testName: "it should return true if given 'madam' input",
			args: args{
				str: "madam",
			},
			expect: true,
		},
	}

	for _, tt := range tests {
		if got := checkPalindrome(tt.args.str); got != tt.expect {
			t.Errorf("FAILED!!!: checkPalindrome expect %t to be PASSED but got %t instead", tt.expect, got)
		}
	}
}

func Benchmark_checkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkPalindrome("cosmic")
	}
}
