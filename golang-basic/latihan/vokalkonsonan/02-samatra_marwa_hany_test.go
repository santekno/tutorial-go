package main

import "testing"

func Test_vokalkosnan(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "vokal",
			args: args{
				input: "a"
			},
			want: "vokal",
		},
		{
			name: "should be consonant",
			args: args{
				letter: "i",
			},
			want: "consonant",
		},
		{
			name: "should be consonant",
			args: args{
				letter: "u",
			},
			want: "consonant",
		},
		{
			name: "should be consonant",
			args: args{
				letter: "e",
			},
			want: "consonant",
		},
		{
			name: "should be consonant",
			args: args{
				letter: "o",
			},
			want: "consonant",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vokalkosnan(tt.args.input); got != tt.want {
				t.Errorf("vokalkosnan() = %v, want %v", got, tt.want)
			}
		})
	}
}
