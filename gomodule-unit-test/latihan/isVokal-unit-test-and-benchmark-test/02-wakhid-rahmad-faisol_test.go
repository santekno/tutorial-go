package main

import "testing"

func Test_isVokal(t *testing.T) {
	type args struct {
		character string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "menentukan vokal",
			args: args{
				character: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVokal(tt.args.character); got != tt.want {
				t.Errorf("isVokal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsVokal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isVokal("a")
	}
}
