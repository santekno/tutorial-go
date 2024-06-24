package main

import "testing"

func Test_vokalDetection(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test lebih dari 1 karakter",
			args: args{
				x: "ac",
			},
			want: "Tolong input 1 karakter saja",
		},
		{
			name: "Test bukan menggunakan huruf",
			args: args{
				x: "1",
			},
			want: "Inputan Bukan Huruf, Tolng Input Huruf",
		},
		{
			name: "Test menggunakan konsonan",
			args: args{
				x: "g",
			},
			want: "Konsonan",
		},
		{
			name: "Test menggunakan Vokal",
			args: args{
				x: "a",
			},
			want: "Vokal",
		},
		{
			name: "Test menggunakan vokal tapi huruf besar",
			args: args{
				x: "A",
			},
			want: "Vokal",
		},
		{
			name: "Test menggunakan Konsonan tapi huruf besar",
			args: args{
				x: "G",
			},
			want: "Konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vokalDetection(tt.args.x); got != tt.want {
				t.Errorf("vokalDetection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkVokalDetection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vokalDetection("x")
	}
}
