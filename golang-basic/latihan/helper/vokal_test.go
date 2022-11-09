package helper

import "testing"

func TestIsVokal(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "The input cannot be less than or more than 1",
			args: args{input: "qwerty"},
			want: "Inputan tidak boleh kurang dari atau lebih dari 1",
		},
		{
			name: "Input is vokal",
			args: args{input: "a"},
			want: "a adalah vokal",
		},
		{
			name: "Input is konsonan",
			args: args{input: "k"},
			want: "k adalah konsonan",
		},
		{
			name: "Input is not vokal or konsonan",
			args: args{input: "4"},
			want: "4 bukan vokal ataupun konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsVokal(tt.args.input); got != tt.want {
				t.Errorf("IsVokal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsVokal(b *testing.B) {
	type args struct {
		input string
	}

	benchmarks := []struct {
		name string
		args args
	}{
		{
			name: "The input cannot be less than or more than 1",
			args: args{input: "qwerty"},
		},
		{
			name: "Input is vokal",
			args: args{input: "a"},
		},
		{
			name: "Input is konsonan",
			args: args{input: "k"},
		},
		{
			name: "Input is not vokal or konsonan",
			args: args{input: "4"},
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsVokal(bm.args.input)
			}
		})
	}
}
