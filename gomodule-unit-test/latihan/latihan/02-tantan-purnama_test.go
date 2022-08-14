package latihan

import "testing"

func TestCheck(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check lenght char",
			args: args{
				s: "vvv",
			},
			want: "too many character",
		},
		{
			name: "check konsonan",
			args: args{
				s: "v",
			},
			want: "konsonan",
		},
		{
			name: "check vocal",
			args: args{
				s: "a",
			},
			want: "vocal",
		},
		{
			name: "check not char",
			args: args{
				s: "!",
			},
			want: "not char",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check(tt.args.s); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Check("o")
	}
}
