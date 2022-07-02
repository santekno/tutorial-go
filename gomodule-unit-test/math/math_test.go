package math

import (
	"testing"
)

func TestAverageGenerate(t *testing.T) {
	type args struct {
		xs []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "must 10",
			args: args{
				xs: []float64{10.0, 10.0},
			},
			want: float64(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.xs); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
		}
	}
}

func TestGanjilGenap(t *testing.T) {
	type args struct {
		angka int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case mengeluarkan ganjil",
			args: args{
				angka: 1,
			},
			want: "ganjil",
		},
		{
			name: "test case mengeluarkan genap",
			args: args{
				angka: 2,
			},
			want: "genap",
		},
		{
			name: "test case mengeluarkan -1",
			args: args{
				angka: -1,
			},
			want: "ganjil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GanjilGenap(tt.args.angka); got != tt.want {
				t.Errorf("GanjilGenap() = %v, want %v", got, tt.want)
			}
		})
	}
}
