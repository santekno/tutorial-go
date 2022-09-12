package main

import "testing"

func TestCek_polidrome(t *testing.T) {
	type args struct {
		polidrome string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "true hasil polidrome",
			args: args{
				polidrome: "madam",
			},
			want: "true",
		},
		{
			name: "false hasil polidrome",
			args: args{
				polidrome: "hewan",
			},
			want: "false",
		},
		{
			name: "gagal hasil polidrome",
			args: args{
				polidrome: "a",
			},
			want: "String harus lebih dari 2 huruf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cek_polidrome(tt.args.polidrome); got != tt.want {
				t.Errorf("Cek_polidrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPolidrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cek_polidrome("madam")
	}
}
