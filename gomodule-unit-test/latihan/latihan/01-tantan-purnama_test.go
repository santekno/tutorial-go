package latihan

import (
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		args     args
		wantTemp bool
	}{
		{
			name:     "Check Polindrome true",
			args:     args{s: "madam"},
			wantTemp: true,
		},
		{
			name:     "Check Polindrome false",
			args:     args{s: "dalam"},
			wantTemp: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTemp := Reverse(tt.args.s); gotTemp != tt.wantTemp {
				t.Errorf("Reverse() = %v, want %v", gotTemp, tt.wantTemp)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("madam")
	}
}
