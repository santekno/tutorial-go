package main

import "testing"

func Test_checkLetter(t *testing.T) {
	type args struct {
		letter string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "is vokal",
			args: args{
				letter: "a",
			},
			want: "Vokal",
		},
		{
			name: "is konsonan",
			args: args{
				letter: "b",
			},
			want: "Konsonan",
		},
		{
			name: "is not Alphabet",
			args: args{
				letter: "?",
			},
			want: "Bukan huruf",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkLetter(tt.args.letter); got != tt.want {
				t.Errorf("checkLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
