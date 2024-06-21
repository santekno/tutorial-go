package main

import "testing"

func Polindrom(pol string) bool {
	len := len(pol)
	for i := 0; i < len/2; i++ {
		if pol[i] != pol[len-i-1] {
			return false
		}
	}
	return true
}

func TestPolindrom(t *testing.T) {
	type args struct {
		pol string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "must_be_true",
			args: args{
				pol: "madam",
			},
			want: bool(true),
		},
		{
			name: "must_be_false",
			args: args{
				pol: "cicak",
			},
			want: bool(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Polindrom(tt.args.pol); got != tt.want {
				t.Errorf("Polindrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPolindrom(b *testing.B) {
	kata := "madam"
	kata1 := "cicak"
	for i := 0; i < b.N; i++ {
		Polindrom(kata)
	}
	for i := 0; i < b.N; i++ {
		Polindrom(kata1)
	}
}
