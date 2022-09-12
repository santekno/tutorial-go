package main

import "testing"

func TestCheck_vokal(t *testing.T) {
	type args struct {
		input_string string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success vokal",
			args: args{
				input_string: "a",
			},
			want: "a adalah vokal",
		}, {
			name: "success konsonan",
			args: args{
				input_string: "z",
			},
			want: "z adalah konsonan",
		},
		{
			name: "failed alfabet",
			args: args{
				input_string: "aa",
			},
			want: "alfabet harus 1 huruf",
		}, {
			name: "alfabet kosong",
			args: args{
				input_string: "",
			},
			want: "alfabet tidak boleh kosong",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check_string(tt.args.input_string); got != tt.want {
				t.Errorf("Check_vokal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Check_string("a")
	}
}
