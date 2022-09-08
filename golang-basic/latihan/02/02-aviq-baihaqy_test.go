package main

import (
	"strings"
	"testing"
)

func Test_isVowel(t *testing.T) {
	rh := strings.NewReader("h")
	chrh, _ := rh.ReadByte()

	ra := strings.NewReader("a")
	chra, _ := ra.ReadByte()

	type args struct {
		chr byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "huruh Konsonan",
			args: args{
				chr: chrh,
			},
			want: "h termsuk karakter Konsonan",
		},
		{
			name: "huruf Vokal",
			args: args{
				chr: chra,
			},
			want: "a termsuk karakter Vokal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isVowel(tt.args.chr)
		})
	}
}

func BenchmarkIsVowel(b *testing.B) {
	rh := strings.NewReader("h")
	chrh, _ := rh.ReadByte()

	for i := 0; i < b.N; i++ {
		isVowel(chrh)
	}

}
