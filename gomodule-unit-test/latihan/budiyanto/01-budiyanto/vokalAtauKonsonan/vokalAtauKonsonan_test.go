package main

import (
	"testing"
)

func Test_checkVowelOrConsonant(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case mengeluarkan vokal",
			args: args{
				input: "a",
			},
			want: "Vokal",
		},
		{
			name: "test case mengeluarkan konsonan",
			args: args{
				input: "g",
			},
			want: "Konsonan",
		},
		{
			name: "test case mengeluarkan unknown",
			args: args{
				input: "1",
			},
			want: "Unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkVowelOrConsonant(tt.args.input); got != tt.want {
				t.Errorf("checkVowelOrConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_checkVowelOrConsonant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkVowelOrConsonant("a")
	}
}
