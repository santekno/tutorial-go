package main

import "testing"

func Test_checkConsonant(t *testing.T) {
	type args struct {
		textInput string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case vocal",
			args: args{
				textInput: "a",
			},
			want: "Text input is vocal",
		},
		{
			name: "test case consonant",
			args: args{
				textInput: "b",
			},
			want: "Text input is consonant",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkConsonant(tt.args.textInput); got != tt.want {
				t.Errorf("checkConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPolindrome(t *testing.T) {
	type args struct {
		textInput string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case polindrome",
			args: args{
				textInput: "madam",
			},
			want: "Text input is polindrome",
		},
		{
			name: "test case not polindrome",
			args: args{
				textInput: "polindrome",
			},
			want: "Text input is not polindrome",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPolindrome(tt.args.textInput); got != tt.want {
				t.Errorf("checkPolindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_checkConsonant(b *testing.B) {
	var textInput string = "a"
	for i := 0; i < b.N; i++ {
		checkConsonant(textInput)
	}
}

func Benchmark_checkPolindrome(b *testing.B) {
	var textInput string = "madam"
	for i := 0; i < b.N; i++ {
		checkPolindrome(textInput)
	}
}
