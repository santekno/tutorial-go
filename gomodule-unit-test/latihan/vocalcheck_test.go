package latihantest

import "testing"

func Test_vocalcheck(t *testing.T) {
	type args struct {
		dicek string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Check Vocal A ",
			args: args{
				dicek: "a",
			},
			want: "huruf Vokal",
		},
		{
			name: "Check Konsonan T ",
			args: args{
				dicek: "t",
			},
			want: "huruf Konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vocalcheck(tt.args.dicek); got != tt.want {
				t.Errorf("vocalcheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkVocalcheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vocalcheck("a")
	}
}
