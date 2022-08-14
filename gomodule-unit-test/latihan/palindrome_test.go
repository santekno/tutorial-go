package latihantest

import "testing"

func Test_cekpalindrom(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Check Palindrom kapak",
			args: args{
				input: "kapak",
			},
			want: true,
		},
		{
			name: "Check Palindrom mobil",
			args: args{
				input: "mobil",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cekpalindrom(tt.args.input); got != tt.want {
				t.Errorf("cekpalindrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkCekpalindrom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cekpalindrom("kapak")
	}
}
