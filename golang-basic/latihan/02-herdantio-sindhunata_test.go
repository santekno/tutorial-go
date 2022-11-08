package main

import "testing"

func Test_isVocal(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "SUCCESS on input \"i\" is vocal",
			args: args{
				input: "i",
			},
			want: true,
		},
		{
			name: "SUCCESS on input \"katak\" is consonant",
			args: args{
				input: "k",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVocal(tt.args.input); got != tt.want {
				t.Errorf("isVocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsVocal(b *testing.B) {
	vocalString := "a"

	for i := 0; i < b.N; i++ {
		isVocal(vocalString)
	}
}
