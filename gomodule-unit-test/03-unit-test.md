# Unit Test

## Testing Menggunakan Library Go
Pemrograman tidak mudah bahkan `programmer` terbaik tidak mampu menulis program yang bekerja persis seperti yang diinginkan setiap saat. Oleh karena itu bagian penting dari proses pengembangan perangkat lunak adalah pengujian (`testing`). Menulis `test` untuk kode kita adalah cara yang baik untuk memastikan kualitas dan meningkatkan keandalan.

Go menyediakan package testing, berisikan banyak sekali tools untuk keperluan unit test. Pada chapter ini kita akan belajar mengenai testing, benchmark, dan juga testing menggunakan testify.

Go menyertakan program khusus yang membuat tes menulis lebih mudah, jadi mari kita buat beberapa `test` untuk paket yang kita buat di sesi ini. Di folder `01-math` buat file baru bernama `math_test.go` yang berisi ini.

Sebelumnya kita telah buat suatu function sebagai berikut ini.
```go
package main

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
```
agar kita bisa membuat unit test maka kita generate menggunakan `vscode`, maka akan membuat file baru `math_test.go` dan men-generate fungsi `TestAverage` dengan isi dibawah ini.
```go
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
```

### Unit Test Program Ganjil Genap
Buat fungsi untuk menentukan bahwa program tersebut mengeluarkan informasi yang diinput terkait angka `ganjil` dan `genap`. 
```go
func GanjilGenap(angka int) string {
	if angka%2 == 0 {
		return "genap"
	}
	return "ganjil"
}
``` 
Maka kita akan buat unit test-nya seperti ini:
```go
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
```

### Testing Tabung
Pertama siapkan terlebih dahulu sebuah struct `Tabung`. Variabel object hasil struct ini nantinya kita gunakan sebagai bahan testing.

```go
package math

import "math"

type Tabung struct {
	Jarijari, Tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Phi * math.Pow(t.Jarijari, 2) * t.Tinggi
}

func (t Tabung) Luas() float64 {
	return 2 * math.Phi * t.Jarijari * (t.Jarijari + t.Tinggi)
}

func (t Tabung) KelilingAlas() float64 {
	return 2 * math.Phi * t.Jarijari
}
```
Maka kita akan membuat unit test satu-satu dari fungsi yang sudah kita buat. Bisa dilihat sebagai berikut.
* Unit Test untuk fungsi Volume
```go
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
			want: float64(792.8366544874485),
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
```
* Unit Test untuk fungsi Luas Permukaan
```go
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
			want: float64(385.092089322475),
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
```
* Unit test untuk Keliling Alas
```go
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
			want: float64(22.65247584249853),
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
```
Cara eksekusi testing adalah menggunakan command `go test`. Argument `-v` atau verbose digunakan menampilkan semua output log pada saat pengujian.

Jalankan program seperti di bawah ini, terlihat bahwa tidak ada test yang fail.
```bash
➜  tabung git:(main) ✗ go test -v
=== RUN   TestTabung_Volume
=== RUN   TestTabung_Volume/testing_hitung_volume
--- PASS: TestTabung_Volume (0.00s)
    --- PASS: TestTabung_Volume/testing_hitung_volume (0.00s)
=== RUN   TestTabung_Luas
=== RUN   TestTabung_Luas/testing_hitung_luas_permukaan
--- PASS: TestTabung_Luas (0.00s)
    --- PASS: TestTabung_Luas/testing_hitung_luas_permukaan (0.00s)
=== RUN   TestTabung_KelilingAlas
=== RUN   TestTabung_KelilingAlas/testing_hitung_keliling_alas
--- PASS: TestTabung_KelilingAlas (0.00s)
    --- PASS: TestTabung_KelilingAlas/testing_hitung_keliling_alas (0.00s)
PASS
ok      github.com/santekno/tabung      0.486s
```

### Method Test
| Method | Kegunaan |
| :- | :--------- |
| `Log()` |	Menampilkan log |
| `Logf()` |	Menampilkan log menggunakan format |
| `Fail()` |	Menandakan terjadi Fail() dan proses testing fungsi tetap diteruskan |
| `FailNow()` |	Menandakan terjadi Fail() dan proses testing fungsi dihentikan |
| `Failed()` |	Menampilkan laporan fail |
| `Error()` |	Log() diikuti dengan Fail() |
| `Errorf()` |	Logf() diikuti dengan Fail() |
| `Fatal()` |	Log() diikuti dengan failNow() |
| `Fatalf()` |	Logf() diikuti dengan failNow() |
| `Skip()` |	Log() diikuti dengan SkipNow() |
| `Skipf()` |	Logf() diikuti dengan SkipNow() |
| `SkipNow()` |	Menghentikan proses testing fungsi, dilanjutkan ke testing fungsi  |setelahnya
| `Skiped()` |	Menampilkan laporan skip |
| `Parallel()` |	Menge-set bahwa eksekusi testing adalah parallel |

## Benchmark

Package `testing` selain berisikan tools untuk testing juga berisikan tools untuk benchmarking. Cara pembuatan benchmark sendiri cukup mudah yaitu dengan membuat fungsi yang namanya diawali dengan `Benchmark` dan parameternya bertipe `*testing.B`.

Sebagai contoh, kita akan mengetes performa perhitungan luas tabung. Siapkan fungsi dengan nama `BenchmarkHitungLuas()` dengan isi adalah kode berikut.
```go
func BenchmarkHitungLuas(b *testing.B) {
	tabung := Tabung{Jarijari: 7, Tinggi: 10}
	for i := 0; i < b.N; i++ {
		tabung.Luas()
	}
}
```
Jalankan test menggunakan argument `-bench=.`, argumen ini digunakan untuk menandai bahwa selain `testing` terdapat juga `benchmark` yang perlu diuji.
```bash
➜  tabung git:(main) ✗ go test -v -bench=.              
=== RUN   TestTabung_Volume
=== RUN   TestTabung_Volume/testing_hitung_volume
--- PASS: TestTabung_Volume (0.00s)
    --- PASS: TestTabung_Volume/testing_hitung_volume (0.00s)
=== RUN   TestTabung_Luas
=== RUN   TestTabung_Luas/testing_hitung_luas_permukaan
--- PASS: TestTabung_Luas (0.00s)
    --- PASS: TestTabung_Luas/testing_hitung_luas_permukaan (0.00s)
=== RUN   TestTabung_KelilingAlas
=== RUN   TestTabung_KelilingAlas/testing_hitung_keliling_alas
--- PASS: TestTabung_KelilingAlas (0.00s)
    --- PASS: TestTabung_KelilingAlas/testing_hitung_keliling_alas (0.00s)
goos: darwin
goarch: arm64
pkg: github.com/santekno/tabung
BenchmarkHitungLuas
BenchmarkHitungLuas-8           1000000000               0.3317 ns/op
PASS
ok      github.com/santekno/tabung      2.557s
```

Arti dari 1000000000 0.3317 ns/op adalah, fungsi di atas di-test sebanyak 1 miliar kali, hasilnya membutuhkan waktu rata-rata 0.3317 nano detik untuk run satu fungsi.


## Testify

Package testify berisikan banyak sekali tools yang bisa dimanfaatkan untuk keperluan testing di Go.

Testify bisa di-download pada [github.com/stretchr/testify](github.com/stretchr/testify) menggunakan go get.

Di dalam `testify` terdapat 5 package dengan kegunaan berbeda-beda satu dengan lainnya. Detailnya bisa dilihat pada tabel berikut.

| Package | Kegunaan |
| :- | :--------- |
|`assert` |	Berisikan tools standar untuk testing |
|`http` |	Berisikan tools untuk keperluan testing http |
|`mock` |	Berisikan tools untuk mocking object |
|`require` |	Sama seperti assert, hanya saja jika terjadi fail pada saat test akan  |menghentikan eksekusi program
|`suite` |	Berisikan tools testing yang berhubungan dengan struct dan method |

Pada sesi ini akan kita contohkan bagaimana penggunaan package `assert`. Silakan perhatikan contoh berikut.
```bash
➜  tabung git:(main) ✗ go test -v         
=== RUN   TestHitungVolume
--- PASS: TestHitungVolume (0.00s)
=== RUN   TestHitungLuas
--- PASS: TestHitungLuas (0.00s)
=== RUN   TestHitungKeliling
--- PASS: TestHitungKeliling (0.00s)
PASS
ok      github.com/santekno/tabung      0.326s
```