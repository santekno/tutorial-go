package main

import (
	"testing"
)

func Test_isVocal(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should be Vokal",
			args: args{
				str: "A",
			},
			want: "Vokal",
		},
		{
			name: "should be Konsonan",
			args: args{
				str: "h",
			},
			want: "Konsonan",
		},
		{
			name: "should be error, input can't be number",
			args: args{
				str: "1",
			},
			wantErr: true,
		},
		{
			name: "should be error, max 1 character",
			args: args{
				str: "Achmad",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isVocal(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("isVocal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isVocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
