package main

import "testing"

func Test_isVocal(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Success pada input karakter i",
			args: args{
				input: "i",
			},
			want: true,
		},
		{
			name: "Failed pada input karakter k",
			args: args{
				input: "k",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVocal(tt.args.input); got != tt.want {
				t.Errorf("isVocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
