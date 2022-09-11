package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "empty string",
			args: args{
				input: "",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ONE LETTER",
			args: args{
				input: "a",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "WITH NUMBER",
			args: args{
				input: "b1",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isPalindrome(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("isPalindrome() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
