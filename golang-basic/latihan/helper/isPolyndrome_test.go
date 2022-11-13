package helper

import "testing"

func TestIsPolyndrome(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Success Response TestIsPolyndrome",
			args: args{
				word: "katak",
			},
			want: true,
		},
		{
			name: "Failed Response TestIsPolyndrome",
			args: args{
				word: "abcd",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPolyndrome(tt.args.word); got != tt.want {
				t.Errorf("IsPolyndrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPolyndrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPolyndrome("katak")
	}
}
