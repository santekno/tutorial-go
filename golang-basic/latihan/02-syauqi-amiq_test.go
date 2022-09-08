package main

import "testing"

func Test_checkConsonant(t *testing.T) {
	type args struct {
		char string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should be consonant",
			args: args{
				char: "a",
			},
			want: false,
		},
		{
			name: "should be consonant",
			args: args{
				char: "i",
			},
			want: false,
		},
		{
			name: "should be consonant",
			args: args{
				char: "u",
			},
			want: false,
		},
		{
			name: "should be consonant",
			args: args{
				char: "e",
			},
			want: false,
		},
		{
			name: "should be consonant",
			args: args{
				char: "o",
			},
			want: false,
		},
		{
			name: "should hit default value",
			args: args{
				char: "b",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkConsonant(tt.args.char); got != tt.want {
				t.Errorf("checkConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}
