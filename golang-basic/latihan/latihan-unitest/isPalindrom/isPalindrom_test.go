package isPalindrom

import "testing"

func TestIsPalindrom(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should be isPalindrom = true",
			args: args{
				str: "katak",
			},
			want: true,
		},
		{
			name: "should be isPalindrom = false",
			args: args{
				str: "golang",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrom(tt.args.str); got != tt.want {
				t.Errorf("IsPalindrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
