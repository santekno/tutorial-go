package vocalconsonant

import (
	"testing"
)

func TestCheckChar(t *testing.T) {
	type args struct {
		char string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check Vokal",
			args: args{
				char: "a",
			},
			want: "Vokal",
		},
		{
			name: "check Konsonan",
			args: args{
				char: "b",
			},
			want: "Konsonan",
		},
		{
			name: "check karakter lebih dari 1",
			args: args{
				char: "ab",
			},
			want: "Hanya boleh 1 karakter",
		},
		{
			name: "check hanya boleh huruf",
			args: args{
				char: "1",
			},
			want: "Hanya boleh huruf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckChar(tt.args.char); got != tt.want {
				t.Errorf("CheckChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
