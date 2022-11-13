package main

import "testing"

func Test_consonant(t *testing.T) {
	type args struct {
		inp string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "vokal",
			args: args{"a"},
			want: false,
		},
		{
			name: "konsonan",
			args: args{"h"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := consonant(tt.args.inp); got != tt.want {
				t.Errorf("consonant() = %v, want %v", got, tt.want)
			}
		})
	}
}
