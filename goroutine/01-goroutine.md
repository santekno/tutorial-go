# Go-routine

Penggunaan goroutine saat proses yang akan dieksekusi sebagai goroutine harus dibungkus kedalam fungsi. Pada saat pemanggilan fungsi tersebut ditambahkan di depannya dengan perintah `go`. Dengan demikian proses tersebut akan dideteksi sebagai goroutine baru.

Berikut ini contoh sederhana dari penerapan goroutine yang sederhana. Program ini menampilkan 10 baris teks, yang mana 5 baris dicetak menggunakan fungsi biasa dan 5 lainnya dicetak menggunakan goroutine baru.
```go
package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}
func main() {
	runtime.GOMAXPROCS(2)
	go print(5, "halo")
	print(5, "apa kabar")
	var input string
	fmt.Scanln(&input)
}
```

Pada kode di atas, `runtime.GOMAXPROCS(2)` digunakan untuk menentukan jumlah prosesor yang aktif.

Pembuatan goroutine baru ditandai dengan `go`. Contoh pada perintah `go print(5, "halo")`, pada fungsi `print()` dieksekusi sebagai goroutine baru. Fungsi `fmt.Scanln()` digunakan agar proses jalannya applikasi berhenti di baris tersebut *blockig* hingga user menekan tombol `enter`. 

Hal ini diperlukan karena ada kemungkinan proses goroutine `print()` dieksekusi lebih lama dibandingkan waktu selesainya goroutine utama `main()`, mengingat proses keduanya tersebut dijalankan secara `asynchronous`. Jika kita tidak menggunakan perintah tersebut maka, goroutine yang belum selesai secara paksa dihentikan prosesnya karena goroutine utama sudah selesai dijalankan.

```bash
➜  sederhana git:(main) ✗ go run main.go
1 apa kabar
2 apa kabar
3 apa kabar
4 apa kabar
5 apa kabar
1 halo
2 halo
3 halo
4 halo
5 halo

➜  sederhana git:(main) ✗ go run main.go
1 halo
2 halo
3 halo
4 halo
5 halo
1 apa kabar
2 apa kabar
3 apa kabar
4 apa kabar
5 apa kabar
```

Tulisan "halo" dan "apa kabar" bermunculan selang-seling disebabkan karena statement `print(5,"halo")` dijalankan dengan goroutine baru, menjadikan tidak saling tunggu dengan perintah `print(5,"apa kabar").

Pada hasil dijalankan 2 kali program diatas. Hasil eksekusi pertama berbeda dengan yang kedua dikarenakan kita menggunakan 2 prosesor. Goroutine mana yang akan dieksekusi terlebih dahulu tergantung kedua prosesor tersebut.

## Penggunaan Fungsi `runtime.GOMAXPROCS()`

Fungsi ini digunakan untuk menentukan jumlah prosesor yang digunakan dalam eksekusi program.

Jumlah prosesor yang diinputkan secara otomatis akan disesuaikan dengan jumlah asli logical processor yang ada pada server atau komputer. Jika angka yg diinputkan adalah lebih, maka dianggap menggunakan semua prosesor yang ada.

## Penggunaan Fungsi fmt.Scanln()
Fungsi ini akan meng-capture semua karakter sebelum *user* menekan tombol `enter`, lalu menyimpannya pada variabel.
```go
func Scanln(a ...interface{}) (n int, err error)
```

Fungsi diatas merupakan skema dari fungsi `fmt.Scanln()`. Fungsi tersebut bisa menampung parameter bertipe `interface{}` dengan jumlah tak terbatas. Tiap parameter akan ditampung di dalam variabel yang dipisah dengan tanda spasi. Agar lebih jelas, silahkan perhatikan contoh dibawah ini.
```go
var s1, s2, s3 string
fmt.Scanln(&s1, &s2, &s3)

// user inputs: "ihsan san tekno"

fmt.Println(s1) // ihsan
fmt.Println(s2) // san
fmt.Println(s3) // tekno
```

Bisa dilihat pada kode di atas, untuk menampung parameter `ihsan san tekno` kita membutuhkan 3 buah variabel. Juga perlu diperhatikan bahwa yang disisipkan sebagai parameter pada pemanggilan fungsi `fmt.Scanln()` adalah referensi variabel, bukan nilai aslinya.