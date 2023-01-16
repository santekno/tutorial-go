package main

import (
	"strings"
	"testing"
)

func isKonsonan(h string) string {
	kons := "bcdfghjklmnpqrstvwxyz"
	vokal := "aiueo"
	if strings.ContainsAny(h, vokal) {
		return "vokal"
	} else if strings.ContainsAny(h, kons) {
		return "konsonan"
	} else {
		return "masukan huruf konsonan atau vokal"
	}
}

func Test_isKonsonan(t *testing.T) {
	type args struct {
		h string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "must_vokal",
			args: args{
				h: "a",
			},
			want: string("vokal"),
		},
		{
			name: "must_konsonan",
			args: args{
				h: "b",
			},
			want: string("konsonan"),
		},
		{
			name: "must_vokal_or_konsonan",
			args: args{
				h: "1",
			},
			want: string("masukan huruf konsonan atau vokal"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isKonsonan(tt.args.h); got != tt.want {
				t.Errorf("isKonsonan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsKonsonan(b *testing.B) {
	huruf := "a"
	for i := 0; i < b.N; i++ {
		isKonsonan(huruf)
	}
}
