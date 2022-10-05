# Mengenal Library String

Sebagai programmer yang sedang belajar Go lang kita perlu mengetahui juga library package bawaan `default` dari Go agar kita lebih paham dan mudah saat nanti berinteraksi dengan package core. Banyak sekali `library` bawaan dari golang yang kita bisa mempermudah kita saat ngoding. Misalkan kita perlu konversi string maka library Go secara default sudah disediakan.

Untuk lebih lanjutnya kita akan bahas satu persatu library yang menurut kita ini perlu dipelajari untuk menunjang kita saat ngoding.

## String
Go menyertakan beberapa fungsi yang bekerja dengan pengolahan `string` diantaranya sebagai contoh berikut:
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		// true
		strings.Contains("test", "es"),
		// 2
		strings.Count("test", "t"),
		// true
		strings.HasPrefix("test", "te"),
		// true
		strings.HasSuffix("test", "st"),
		// 1
		strings.Index("test", "e"),
		// "a-b-c"
		strings.Join([]string{"a", "b","c"}, "-"),
		// == "aaaaa"
		strings.Repeat("a", 5),
		// "bbaa"
		strings.Replace("aaaa", "a", "b", 2),
		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"),
		// "test"
		strings.ToLower("TEST"),
		// "TEST"
		strings.ToUpper("test"),
	)
}
```

Bisa kita lihat beberapa fungsi yang terdapat pada library `string` yang bisa kita gunakan untuk beberapa operational. 
* Fungsi `Contains` ini sangat berguna sekali jika kita sedang mencari data yang di dalam satu variabel itu berisi data yang ingin kita seleksi. 
* Fungsi `Count` digunakan untuk menghitung karakter yang terseleksi pada kalimat atau kata.
* Fungsi `HasPrefix` juga bisa kita pakai untuk pendeteksi jika kata terdapat huruf pertamya `te` maka akan lebih mudah menggunakan fungsi ini.
* Fungsi `HasSuffix` digunakan untuk mendeteksi apakah memiliki akhiran yang huruf yang sesuai dengan yang kita inputkan.
* Fungsi `Index` berguna untuk mengetahui karakter yang kita cari itu berada pada index ke berapa pada kalimat atau kata.
* Fungsi `Join` digunakan untuk menggabungkan karakter huruf array menjadi satu kalimat / kata dengan ditambahkan delimiter.
* Fungsi `Repeat` digunakan untuk perulangan sebesar input yg kita inginkan.
* Fungsi `Replace` digunakan untuk menimpa suatu karakter/kata dari kata yang sudah kita tentukan.
* Fungsi `Split` digunakan untuk mengambil data string dari karakter dengan delimiter tertentu.
* Fungsi `ToLower` digunakan untuk kata / karakter agar menjadi huruf kecil semua.
* Fungsi `ToUpper` digunakan untuk kata / karakter agar menjadi kapital.

Memudahkan sekali bukan? Golang telah menyediakan fungsi-fungsi pendukung agar kita lebih mudah lagi untuk mengoperasikannya.

Terkadang kita perlu menggunakan string sebagai data biner. Untuk mengonversi string menjadi sepotong byte (dan sebaliknya) itu bisa dilakukan seperti ini
```go
arr := []byte("test")
str := string([]byte{'t','e','s','t'})
```

Dan masih banyak lagi beberapa fungsi pendukung untuk memanipulasi `string` yang sudah disediakan golang. Bisa dilihat juga Library String [disini](https://blog.golang.org/strings).