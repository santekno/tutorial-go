package pkg

import "testing"

func TestIsVocal(t *testing.T) {
	type args struct {
		char string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success expected vocal",
			args: args{
				char: "a",
			},
			want: "Vocal",
		},
		{
			name: "success expected vocal",
			args: args{
				char: "i",
			},
			want: "Vocal",
		},
		{
			name: "success expected consonant",
			args: args{
				char: "b",
			},
			want: "Consonant",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsVocal(tt.args.char); got != tt.want {
				t.Errorf("IsVocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsVocal(b *testing.B) {
	char := "a"
	for i := 0; i < b.N; i++ {
		IsVocal(char)
	}
}
