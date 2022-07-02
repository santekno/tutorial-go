# Control Structures

## For
Perulangan pada golang hanya operasi `for` saja tidak punya `foreach`, `while`, `do`, `until` tetapi dengan menggunakan `for` ini semua operasi dalam aritmatika ataupun algoritma sudah bisa dilakukan. Misalkan contoh dasar kita akan mengeluarkan bilangan/angka dari 0-1. Maka code-nya seperti ini
```go
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
```
Maka hasil dari program diatas mengeluarkan output
```bash
$ go run golang-basic/06-structure/perulangan-angka.go
1
2
3
4
5
6
7
8
9
10
```


## If ... else ...
Kondisional pada umumnya biasanya menggunakan `if ... else`. Dan hematnya code golang itu tidak perlu ada kurung tutup dan kurung buka yang biasanya program yang lain harus menggunakan itu. Contoh program kondisional seperti ini
```go
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " Genap")
		} else {
			fmt.Println(i, " Ganjil")
		}
	}
}
```

## Switch
Kondisional program juga bisa menggunakan `switch` yang mana lebih mudah untuk satuan pengecekan. Bila lebih complex suatu kondisi program lebih bisa dipahami menggunakan `if ... else ...` saja.
Berikut ini contoh untuk `switch` pada program untuk mengeluarkan nama indonesia dalam bahasa.
```go
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		switch i {
		case 1:
			fmt.Println(i, " Satu")
		case 2:
			fmt.Println(i, " Dua")
		case 3:
			fmt.Println(i, " Tiga")
		case 4:
			fmt.Println(i, " Empat")
		case 5:
			fmt.Println(i, " Lima")
		case 6:
			fmt.Println(i, " Enam")
		case 7:
			fmt.Println(i, " Tujuh")
		case 8:
			fmt.Println(i, " Delapan")
		case 9:
			fmt.Println(i, " Sembilan")
		case 10:
			fmt.Println(i, " Sepuluh")
		default:
			fmt.Println(i, " Tidak terdeteksi")
		}
	}
}
```

## Arrays
Penampung data yang biasanya berurutan dengan penampung yang sudah ditentukan. Biasanya jika ingin mengisi data `array` kita perlu definisikan terlebih dahulu berapa data yang akan ditampun dalam suatu array tersebut.

Misalkan kita tulis program seperti dibawah ini
```go
  var x [5]int
```
Artinya, kita menyediakan penyimpanan berurutan (sequence) yang berisi 5 data yang bisa diisi oleh program. Misalkan kita akan isi variable `x` tersebut pada posisi terakhir (index ke-4). Maka program akan seperti ini.
```go
package main

import "fmt"

func main() {
     var x [5]int
     x[4] = 100
     fmt.Println(x)
}
```
maka output dari program tersebut data yang ke-5 akan berisi 100.
```bash
$ go run golang-basic/06-structure/array-structures.go 
[0 0 0 0 100]
```
Definisikan terlebih dahulu berapa yang akan disimpan element-nya lalu kita simpan data tersebut dalam salah satu element ke-`index`
> Ingat dalam array itu tiap element akan dimulai dengan index ke-0

### Contoh Program 2
Kita akan menjumlakan semua element dari suatu `array`. Maka program-nya seperti ini
```go
func main() {
     var x [5]float64
     x[0] = 98
     x[1] = 93
     x[2] = 77
     x[3] = 82
     x[4] = 83
     var total float64 = 0
     for i := 0; i < 5; i++ {
           total += x[i]
     }
     fmt.Println(total)
}
```
Deklarasi dan pengisian array bisa juga sekaligus didefinisikan atau lebih ringkasnya lihat dibawah ini:
```go
  x := [5]float64{ 98, 93, 77, 82, 83 }
```

## Slice
Slice adalah segmen dari array. Seperti array, irisan dapat diindeks dan memiliki panjang. Tidak seperti array, panjang ini diizinkan untuk berubah. Berikut adalah contoh program:
```go
  var x []float64
```
atau bisa juga menggunakan `make`
```go
  x := make([]float64, 5)
```
Slice ini bisa kita gunakan untuk array yang bisa kita tambah kurang sesuai dengan kebutuhan kita. misalkan kita definisikan 5 tetapi ingin menambahkan lagi dibelakang data yang ke 5 itu, maka lebih mudah menggunakan slice. Contoh program dibawah ini.
```go
package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}
```
lalu bisa juga kita menggunakan `copy`, coba kita ubah program diatas menjadi seperti ini.
```go
func main() {
     slice1 := []int{1,2,3}
     slice2 := make([]int, 2)
     copy(slice2, slice1)
     fmt.Println(slice1, slice2)
}
```
Maka apa yang akan terjadi? Hanya angka `1 dan 2` saja yang ada pada `slice2`.

## Maps
Struktur data yang tidak berurutan sehingga kita bebas menyimpan-nya dalam posisi mana saja. Berikut ini contoh definisi map pada variable `x`.
```go
  var x map[string]int
```
Misalkan kita akan membuat program seperti dibawah ini.
```go
  var x map[string]int
  x["key"] = 10
  fmt.Println(x)
```
maka akan mengeluarkan error `panic` seperti ini
```bash
$ go run golang-basic/06-structure/error-panic-map.go 
panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()
        /Users/ihsanarif/Documents/ihsan/tutorial/tutorial-go/golang-basic/06-structure/error-panic-map.go:7 +0x44
exit status 2
```
maka harus kita inisialisasi terlebih dahulu
```go
  x := make(map[string]int)
  x["key"] = 10
  fmt.Println(x["key"])
```
Jika kita ingin menghapus salah satu dari maps tersebut kita bisa menggunakan func `delete`.
```go
delete(x, 1)
```
Contoh program kita akan menyimpan data unsur-unsur dalam kita menggunakan maps.
```go
package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Ho"])
}
```
Apa yang terjadi jika kita memanggil element yang tidak ada pada map? Akan terjadi panic atau program mengeluarkan kosong?
Ya, khusus map secara `default`, data tersebut jika tidak terdapat pada suatu index map maka akan menghasilkan output kosong `""`.

Trik agar kita mengetahui bahwa dalam `maps` tersebut terdapat index `x` atau tidak kita perlu penambahan kondisi seperti ini.
```go
  name, ok := elements["Un"]
  fmt.Println(name, ok)
```

Dalam maps ini juga kita bisa buat multiple map dalam satu variable. Bisa dilihat contoh program seperti dibawah ini.
```go
package main

import "fmt"

func main() {
	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
```
## Latihan
* Modifikasi program menjumlahkan semua element lalu tampilkan jumlah, median, mean, modus

