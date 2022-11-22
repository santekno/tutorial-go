package main

import "testing"

func Test_isVokal(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "has vocal letter, must return true",
			args: args{
				input: "a",
			},
			want: true,
		},
		{
			name: "has consonant letter, must return false",
			args: args{
				input: "b",
			},
			want: false,
		},
		{
			name: "input more than 1 letter, must return false",
			args: args{
				input: "ba",
			},
			want: false,
		},
		{
			name: "input not letter, must return false",
			args: args{
				input: "!",
			},
			want: false,
		},
		{
			name: "input is number, must return false",
			args: args{
				input: "1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVokal(tt.args.input); got != tt.want {
				t.Errorf("isVokal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsVokal(b *testing.B) {
	input := "a"
	for i := 0; i < b.N; i++ {
		isVokal(input)
	}
}
