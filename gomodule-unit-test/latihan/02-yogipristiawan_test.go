package latihan

import (
	"errors"
	"reflect"
	"testing"
)

func Test_validateInput(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		testName string
		args     args
		expect   error
	}{
		{
			testName: "it should return error if given input more than 1 character long",
			args: args{
				str: "qwerty",
			},
			expect: errors.New("ERROR: input must be 1 character long"),
		},
		{
			testName: "it should return error if given input is not alphabet character",
			args: args{
				str: "0",
			},
			expect: errors.New("ERROR: input must be an alphabet"),
		},
		{
			testName: "it should return error nil if given input is valid alphabet character",
			args: args{
				str: "a",
			},
			expect: nil,
		},
	}

	for _, tt := range tests {
		if got := validateInput(tt.args.str); reflect.TypeOf(got) != reflect.TypeOf(tt.expect) {
			t.Errorf("FAILED: validateInput expect %T to be PASSED but got %T instead", tt.expect, got)
		}
	}
}

func Benchmark_validateInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validateInput("a")
	}
}

func Test_checkVowel(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		testName string
		args     []args
		expect   bool
	}{
		{
			testName: "it should return true if given input is vowel",
			args: []args{
				{
					str: "a",
				},
				{
					str: "A",
				},
			},
			expect: true,
		},
		{
			testName: "it should return false if given input is not vowel",
			args: []args{
				{
					str: "p",
				},
				{
					str: "P",
				},
			},
			expect: false,
		},
	}

	for _, tt := range tests {
		for _, val := range tt.args {
			if got := checkVowel(val.str); got != tt.expect {
				t.Errorf("FAILED!!!: checkVowel expect %t to be PASSED but got %t instead", tt.expect, got)
			}
		}
	}
}

func Benchmark_checkVowel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkVowel("a")
	}
}
