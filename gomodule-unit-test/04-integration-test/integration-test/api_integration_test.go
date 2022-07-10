//go:build integration

package integrationtest

import (
	"net/http"
	"testing"
)

func TestMathClient_APISum(t *testing.T) {
	type fields struct {
		Token string
		Host  string
	}
	type args struct {
		i string
		j string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     int
		wantErr  bool
		wantCode int
	}{
		{
			name: "case sukses jumlah data",
			fields: fields{
				Token: "abcdef123",
				Host:  "localhost:9000",
			},
			args: args{
				i: "2",
				j: "2",
			},
			want:     4,
			wantErr:  false,
			wantCode: http.StatusOK,
		},
		{
			name: "token tidak diset",
			fields: fields{
				Token: "abc",
				Host:  "localhost:9000",
			},
			args: args{
				i: "2",
				j: "2",
			},
			want:     0,
			wantErr:  true,
			wantCode: http.StatusUnauthorized,
		},
		{
			name: "host tidak sesuai",
			fields: fields{
				Token: "abcdef123",
				Host:  "localhost:500",
			},
			args: args{
				i: "2",
				j: "2",
			},
			want:     0,
			wantErr:  true,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "case parameter kosong",
			fields: fields{
				Token: "abcdef123",
				Host:  "localhost:9000",
			},
			args: args{
				i: "",
				j: "",
			},
			want:     0,
			wantErr:  true,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "case parameter bukan integer",
			fields: fields{
				Token: "abcdef123",
				Host:  "localhost:9000",
			},
			args: args{
				i: "a",
				j: "a",
			},
			want:     0,
			wantErr:  true,
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &MathClient{
				Token: tt.fields.Token,
				Host:  tt.fields.Host,
			}
			got, gotCode, err := c.APISum(tt.args.i, tt.args.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("MathClient.APISum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MathClient.APISum() = %v, want %v", gotCode, tt.wantCode)
			}
			if gotCode != tt.wantCode {
				t.Errorf("MathClient.APISum() = %v, want %v", gotCode, tt.wantCode)
			}
		})
	}
}
