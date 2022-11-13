package main

import "testing"

func Test_consonantVocal(t *testing.T) {
	type args struct {
		character string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "success consonant vocal",
			args: args{
				"a",
			},
			want: "vocal",
		},
		{
			name: "success consonant vocal",
			args: args{
				"b",
			},
			want: "consonant",
		},
		{
			name: "success consonant vocal",
			args: args{
				"@",
			},
			want: "Character must be alphabet",
		},
		{
			name: "success consonant vocal",
			args: args{
				"ab",
			},
			want: "Only one character allowed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := consonantVocal(tt.args.character); got != tt.want {
				t.Errorf("consonantVocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkConsonantVocal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		consonantVocal("a")
	}
}
