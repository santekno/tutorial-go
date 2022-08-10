package latihan

import "testing"

func TestIsVocalOrConsonant(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case return huruf vokal",
			args: args{
				value: "a",
			},
			want: "Vokal",
		},
		{
			name: "Case return huruf konsonan",
			args: args{
				value: "l",
			},
			want: "Konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsVocalOrConsonant(tt.args.value); got != tt.want {
				t.Errorf("IsVocalOrConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsVocalOrConsonant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsVocalOrConsonant("j")
	}
}
