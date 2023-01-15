package main

import "testing"

func Test_polindrone(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test polindrome true",
			args: args{
				x: "madam",
			},
			want: true,
		},
		{
			name: "Test polindrome true",
			args: args{
				x: "kaka",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := polindrone(tt.args.x); got != tt.want {
				t.Errorf("polindrone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPolindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		polindrone("madam")
	}
}
