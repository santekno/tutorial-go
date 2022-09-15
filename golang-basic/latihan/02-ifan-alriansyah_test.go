package main

import "testing"

func Test_isVowelOrConsonant(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Check is Vowel",
			args: args{
				val: "a",
			},
			want: "Vokal",
		},
		{
			name: "Check is Vowel",
			args: args{
				val: "i",
			},
			want: "Vokal",
		},
		{
			name: "Check is Vowel",
			args: args{
				val: "u",
			},
			want: "Vokal",
		},
		{
			name: "Check is Vowel",
			args: args{
				val: "e",
			},
			want: "Vokal",
		},
		{
			name: "Check is Vowel",
			args: args{
				val: "o",
			},
			want: "Vokal",
		},
		{
			name: "Check is Consonan",
			args: args{
				val: "b",
			},
			want: "Konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVowelOrConsonant(tt.args.val); got != tt.want {
				t.Errorf("isVowelOrConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsVowelOrConsonant(t *testing.B) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Check is Vowel",
			args: args{
				val: "a",
			},
			want: "Vokal",
		},
		{
			name: "Check is Vowel",
			args: args{
				val: "i",
			},
			want: "Vokal",
		},
		{
			name: "Check is Vowel",
			args: args{
				val: "u",
			},
			want: "Vokal",
		},
		{
			name: "Check is Vowel",
			args: args{
				val: "e",
			},
			want: "Vokal",
		},
		{
			name: "Check is Vowel",
			args: args{
				val: "o",
			},
			want: "Vokal",
		},
		{
			name: "Check is Consonan",
			args: args{
				val: "b",
			},
			want: "Konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.B) {
			if got := isVowelOrConsonant(tt.args.val); got != tt.want {
				t.Errorf("isVowelOrConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}
