# Pointers
Kita tahu bahwa untuk mengirimkan suatu variable ke dalam fungsi bisa berupa parameter dan pointer. Pointer ini biasanya berupa alokasi data dengan alamat tertentu.
Contoh bisa dilihat pada program dibawah ini.
```go
package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}
```
## Penggunakan * dan &
Dalam Go, sebuah pointer direpresentasikan menggunakan karakter * (tanda bintang) diikuti dengan tipe dari nilai yang disimpan. Dalam fungsi nol xPtr adalah pointer ke int.

`*` juga digunakan untuk "dereference" variabel pointer. Dereferensi pointer memberi kita akses ke nilai yang ditunjuk pointer. Saat kita menulis `*xPtr = 0` kita mengatakan "simpan int 0 di lokasi memori yang dirujuk xPtr". Jika kita mencoba `xPtr = 0` malah kita akan mendapatkan kesalahan kompiler karena xPtr bukan `int` melainkan `*int`, yang hanya bisa diberikan `*int` lain.

Akhirnya kita menggunakan `&` operator untuk menemukan alamat variabel. `&x` mengembalikan `*int` (penunjuk ke int) karena `x` adalah int. Inilah yang memungkinkan kita untuk memodifikasi variabel asli. `&x` di main dan xPtr di nol merujuk ke lokasi memori yang sama.

Cara lain untuk mendapatkan pointer adalah dengan menggunakan fungsi baru bawaan:
```go
package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}
```

Dalam beberapa bahasa pemrograman ada perbedaan yang signifikan antara menggunakan `new` dan `&`, dengan sangat hati-hati diperlukan untuk akhirnya menghapus apa pun yang dibuat dengan `new`. Go tidak seperti itu, karena memiliki garbage collection yaitu memori dibersihkan secara otomatis ketika tidak ada lagi yang merujuk padanya.

## Latihan
Jalankan program dibawah ini.
```go
func square(x *float64) {
    *x = *x * *x
}
func main() {
  x := 1.5
  square(&x) 
}
```
Apa hasil output dari program diatas?