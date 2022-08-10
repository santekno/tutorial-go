package polindrome

import "testing"

func BenchmarkPolindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		polindrome("katak")
	}
}

func BenchmarkPolindromeTableTest(b *testing.B) {
	type args struct {
		kalimat string
	}
	tests := []struct {
		name       string
		parameters args
	}{
		{
			name: "test number 1",
			parameters: args{
				kalimat: "katak",
			},
		},
		{
			name: "test number 2",
			parameters: args{
				kalimat: "bobo",
			},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				polindrome(tt.parameters.kalimat)
			}
		})
	}
}

func Test_polindrome(t *testing.T) {
	type args struct {
		kalimat string
	}
	tests := []struct {
		name       string
		parameters args
		want       bool
	}{
		{
			name: "test number 1",
			parameters: args{
				kalimat: "katak",
			},
			want: true,
		},
		{
			name: "test number 2",
			parameters: args{
				kalimat: "bobo",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := polindrome(tt.parameters.kalimat); got != tt.want {
				t.Errorf("polindrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
