package vokal

import (
	"testing"
)

func TestIsVocal(t *testing.T) {
	type args struct {
		inputChar string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Vocal",
			args: args{
				inputChar: "a",
			},
			want: "Vocal",
		},
		{
			name: "Vocal",
			args: args{
				inputChar: "a",
			},
			want: "Vocal",
		},
		{
			name: "Vocal",
			args: args{
				inputChar: "i",
			},
			want: "Vocal",
		},
		{
			name: "Vocal",
			args: args{
				inputChar: "u",
			},
			want: "Vocal",
		},
		{
			name: "Vocal",
			args: args{
				inputChar: "e",
			},
			want: "Vocal",
		},
		{
			name: "Vocal",
			args: args{
				inputChar: "o",
			},
			want: "Vocal",
		},
		{
			name: "Konsonan",
			args: args{
				inputChar: "b",
			},
			want: "Konsonan",
		},
		{
			name: "Konsonan",
			args: args{
				inputChar: "z",
			},
			want: "Konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsVocal(tt.args.inputChar); got != tt.want {
				t.Errorf("IsVocal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateInput(t *testing.T) {
	type args struct {
		inputChar string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Number",
			args: args{
				inputChar: "2",
			},
			want: false,
		},
		{
			name: "Character",
			args: args{
				inputChar: "@",
			},
			want: false,
		},
		{
			name: "Multiple String",
			args: args{
				inputChar: "aa",
			},
			want: false,
		},
		{
			name: "Single String",
			args: args{
				inputChar: "a",
			},
			want: true,
		},
		{
			name: "String Empty",
			args: args{
				inputChar: "",
			},
			want: false,
		},
		{
			name: "White Space",
			args: args{
				inputChar: " ",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateInput(tt.args.inputChar); got != tt.want {
				t.Errorf("validateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsVocal(b *testing.B) {
	inputChar := "a"
	for i := 0; i < b.N; i++ {
		IsVocal(inputChar)
	}
}
