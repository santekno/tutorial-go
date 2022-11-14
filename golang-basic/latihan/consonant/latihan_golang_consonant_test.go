package main

import "testing"

func Test_isConsonant(t *testing.T) {
	type args struct {
		character string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "must consonant",
			args: args{
				character: string('a'),
			},
			want: false,
		},
		{
			name: "must consonant",
			args: args{
				character: string('b'),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isConsonant(tt.args.character); got != tt.want {
				t.Errorf("isConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}
