package main

import "testing"

func Test_isCharacter(t *testing.T) {
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
			name: "success validate word palindrome",
			args: args{"a"},
			want: true,
		},
		{
			name: "success validate word palindrome",
			args: args{"p"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCharacter(tt.args.character); got != tt.want {
				t.Errorf("isCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
