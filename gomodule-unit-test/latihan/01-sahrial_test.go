package latihan

import "testing"

func TestIsPolindrome(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Cek polindrome kodok",
			args: args{
				value: "kodok",
			},
			want: true,
		},
		{
			name: "Cek polindrome adam",
			args: args{
				value: "adam",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPolindrome(tt.args.value); got != tt.want {
				t.Errorf("IsPolindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsPolindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPolindrome("katak")
	}
}
