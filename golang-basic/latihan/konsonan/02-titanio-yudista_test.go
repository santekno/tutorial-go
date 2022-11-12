package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "a",
		},
		{
			name: "z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_isKonsonanOrVokal(t *testing.T) {
	type args struct {
		substr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Alphabet vokal",
			args: args{
				substr: "E",
			},
			want: "Vokal",
		},
		{
			name: "Alphabet konsonan",
			args: args{
				substr: "S",
			},
			want: "Konsonan",
		},
		{
			name: "Error",
			args: args{
				substr: "1",
			},
			want: "input invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isKonsonanOrVokal(tt.args.substr); got != tt.want {
				t.Errorf("isKonsonanOrVokal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsKonsonanVokal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isKonsonanOrVokal("U")
	}
}
