package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tabung                 = Tabung{Jarijari: 7, Tinggi: 10}
	volumeSeharusnya       = float64(792.8366544874485)
	luasSeharusnya         = float64(385.092089322475)
	kelilingAlasSeharusnya = float64(22.65247584249853)
)

func TestTabung_Volume(t *testing.T) {
	type fields struct {
		Jarijari float64
		Tinggi   float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "testing hitung volume",
			fields: fields{
				Jarijari: 7, Tinggi: 10,
			},
			want: float64(volumeSeharusnya),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tabung{
				Jarijari: tt.fields.Jarijari,
				Tinggi:   tt.fields.Tinggi,
			}
			if got := tr.Volume(); got != tt.want {
				t.Errorf("Tabung.Volume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTabung_Luas(t *testing.T) {
	type fields struct {
		Jarijari float64
		Tinggi   float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "testing hitung luas permukaan",
			fields: fields{
				Jarijari: 7, Tinggi: 10,
			},
			want: float64(luasSeharusnya),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tabung{
				Jarijari: tt.fields.Jarijari,
				Tinggi:   tt.fields.Tinggi,
			}
			if got := tr.Luas(); got != tt.want {
				t.Errorf("Tabung.Luas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTabung_KelilingAlas(t *testing.T) {
	type fields struct {
		Jarijari float64
		Tinggi   float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "testing hitung keliling alas",
			fields: fields{
				Jarijari: 7, Tinggi: 10,
			},
			want: float64(kelilingAlasSeharusnya),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tabung{
				Jarijari: tt.fields.Jarijari,
				Tinggi:   tt.fields.Tinggi,
			}
			if got := tr.KelilingAlas(); got != tt.want {
				t.Errorf("Tabung.KelilingAlas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkHitungLuas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tabung.Luas()
	}
}

func TestHitungVolume(t *testing.T) {
	assert.Equal(t, tabung.Volume(), volumeSeharusnya, "perhitungan volume salah")
}

func TestHitungLuas(t *testing.T) {
	assert.Equal(t, tabung.Luas(), luasSeharusnya, "perhitungan luas salah")
}

func TestHitungKeliling(t *testing.T) {
	assert.Equal(t, tabung.KelilingAlas(), kelilingAlasSeharusnya, "perhitungan keliling salah")
}
