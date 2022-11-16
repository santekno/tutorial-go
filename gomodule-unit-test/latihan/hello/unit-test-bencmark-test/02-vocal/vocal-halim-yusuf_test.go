package main

import "testing"

func Test_isVokalOrKonsonan(t *testing.T) {
	type args struct {
		hur string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "sukses vokal",
			args: args{
				"a",
			},
			want: string("huruf Vokal"),
		},
		{
			name: "sukses konsonan",
			args: args{
				"b",
			},
			want: string("huruf Konsonan"),
		},
		{
			name: "inputan bukan vokal atau konsonan",
			args: args{
				"1",
			},
			want: string("yang Anda masukkan bukan karakter huruf"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVokalOrKonsonan(tt.args.hur); got != tt.want {
				t.Errorf("isVokalOrKonsonan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsVokalOrKonsonan(b *testing.B) {
type args struct {
		hur string
	}
	benchmarks := []struct {
		name string
		args args
	}{
		{
			name: "sukses vokal",
			args: args{
				hur:"a",
			},
		},
		{
			name: "sukses konsonan",
			args: args{
				hur:"b",
			},
		},
		{
			name: "inputan bukan vokal atau konsonan",
			args: args{
				hur:"1",
			},
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isVokalOrKonsonan(bm.args.hur)
			}
		})
	}
}