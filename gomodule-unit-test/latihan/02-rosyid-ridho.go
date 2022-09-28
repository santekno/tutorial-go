package main

import (
	"testing"
)

func Test_isPalindrom(t *testing.T) {
	type args struct {
		kata string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "success palindrom",
			args: args{
				kata: "katak",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrom(tt.args.kata); got != tt.want {
				t.Errorf("isPalindrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isVokal(t *testing.T) {
	type args struct {
		character rune
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success isVocal",
			args: args{
				character: 'a',
			},
		},
		{
			name: "success isVocal",
			args: args{
				character: 'r',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isVokal(tt.args.character)
		})
	}
}
