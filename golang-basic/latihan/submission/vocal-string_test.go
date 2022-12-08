package submission

import (
	"testing"
)

func TestIsVocal(t *testing.T) {
	type args struct {
		character string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "return vocal",
			args: args{
				character: "a",
			},
			want: VOCAL,
		},
		{
			name: "return Kosonan",
			args: args{
				character: "r",
			},
			want: CONSONANT,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsVocal(tt.args.character); got != tt.want {
				t.Errorf("IsVocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_IsVocal(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IsVocal(ramdomString(1))
	}
}
