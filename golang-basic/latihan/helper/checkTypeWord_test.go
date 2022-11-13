package helper

import "testing"

func TestCheckTypeWord(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success Response TestCheckTypeWord Vocal",
			args: args{
				word: "a",
			},
			want: "This is vocal word",
		},
		{
			name: "Success Response TestCheckTypeWord Consonant",
			args: args{
				word: "b",
			},
			want: "This is consonant word",
		},
		{
			name: "Success Response TestCheckTypeWord not Vocal & Consonant",
			args: args{
				word: ".",
			},
			want: "This is not vocal & consonant word",
		},
		{
			name: "Failed Response Because of length Word",
			args: args{
				word: "abcd",
			},
			want: "sorry, you can only input one character",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckTypeWord(tt.args.word); got != tt.want {
				t.Errorf("CheckTypeWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCheckTypeWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckTypeWord("a")
	}
}