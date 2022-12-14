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
		{
			name: "benar palindrome",
			args: args{
				kata: "kasurinirusak",
			},
			want: true,
		}, {
			name: "bukan palindrome",
			args: args{
				kata: "makanenak",
			},
			want: false,
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

func Test_cekKonsonanVokal(t *testing.T) {
	type args struct {
		karakter string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tes konsonan",
			args: args{
				karakter: "k",
			},
			want: 2,
		}, {
			name: "tes vokal",
			args: args{
				karakter: "i",
			},
			want: 1,
		}, {
			name: "bukan huruf",
			args: args{
				karakter: "?",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cekKonsonanVokal(tt.args.karakter); got != tt.want {
				t.Errorf("cekKonsonanVokal() = %v, want %v", got, tt.want)
			}
		})
	}
}
