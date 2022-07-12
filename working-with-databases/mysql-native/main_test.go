package main

import (
	"reflect"
	"testing"
)

func Test_albumsByArtist(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []Album
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := albumsByArtist(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("albumsByArtist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("albumsByArtist() = %v, want %v", got, tt.want)
			}
		})
	}
}
