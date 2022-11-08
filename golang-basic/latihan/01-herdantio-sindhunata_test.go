package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "SUCCESS on input \"katak\" is palindrom",
			args: args{
				input: "katak",
			},
			want: true,
		},
		{
			name: "SUCCESS on input \"adam\" is not palindrom",
			args: args{
				input: "adam",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.input); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
