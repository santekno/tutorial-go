package main

import "testing"

func Test_isVocal(t *testing.T) {
	type args struct {
		vocal string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "konsonan or vocal",
			args: args{
				vocal: "madam",
			},
			want: "Konsonan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
