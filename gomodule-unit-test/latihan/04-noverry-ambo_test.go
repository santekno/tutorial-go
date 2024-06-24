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
			name: "konsonan or vocal",
			args: args{
				input: "madam",
			},
			want: true,
		},
		{
			name: "konsonan or vocal",
			args: args{
				input: "rusak",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
