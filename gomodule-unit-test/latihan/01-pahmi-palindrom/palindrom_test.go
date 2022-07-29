package pahmipalindrom

import "testing"

func TestPalindrom(t *testing.T) {
	type args struct {
		word string
		i    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case false",
			args: args{
				word: "homes",
			},
			want: "false",
		},
		{
			name: "test case true",
			args: args{
				word: "madam",
			},
			want: "true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Palindrom(tt.args.word, tt.args.i); got != tt.want {
				t.Errorf("Palindrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
