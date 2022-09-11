package main

import "testing"

func Test_isConsonant(t *testing.T) {
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
			want:    false,
			wantErr: true,
		},
		{
			name: "input too long",
			args: args{
				input: "asdf",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "input not a letter",
			args: args{
				input: "1",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "input is a vowel",
			args: args{
				input: "a",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "input is a consonant",
			args: args{
				input: "v",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isConsonant(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("isConsonant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}
