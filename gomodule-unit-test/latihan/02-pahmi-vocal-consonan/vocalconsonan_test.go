package main

import "testing"

func TestVocal(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "vocal",
			args: args{
				input: "i",
			},
			want: "vocal",
		},
		{
			name: "Consonan",
			args: args{
				input: "B",
			},
			want: "Consonan",
		},
		{
			name: "simbol",
			args: args{
				input: "}",
			},
			want: "simbol",
		},
		{
			name: "Angka",
			args: args{
				input: "2",
			},
			want: "Angka",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Vocal(tt.args.input); got != tt.want {
				t.Errorf("Vocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTestVocal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Vocal("a")
	}
}
