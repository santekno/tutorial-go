package main

import "testing"

func Test_polindrome(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "success polindrome",
			args: args{
				"wadaw",
			},
			want: true,
		},
		{
			name: "success polindrome",
			args: args{
				"muadz",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := polindrome(tt.args.word); got != tt.want {
				t.Errorf("polindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPolindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		polindrome("wadaw")
	}
}
