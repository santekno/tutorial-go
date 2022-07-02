# Setup Go Module
Pada bagian ini kita akan belajar cara inisialisasi projek menggunakan Go Modules (atau Modules).

## Penjelasan
Go modules merupakan manajemen dependensi resmi untuk Go. Modules ini diperkenalkan pertama kali di `go1.11`, sebelum itu pengembangan projek Go dilakukan dalam `GOPATH`.

Modules digunakan untuk menginisialisasi sebuah projek, sekaligus melakukan manajemen terhadap `3rd party` atau `library` lain yang dipergunakan.

Modules penggunaannya adalah lewat `CLI`. Dan jika kita sudah sukses meng-`install` Go, maka otomatis bisa mempergunakan `Go Modules`.

> `Modules` atau `Module` di sini merupakan istilah untuk project ya. Jadi jangan bingung.

## Inisialisasi
Command `go mod init` ini akan kita gunakan untuk menginisialisasikan projek baru.
Mari kita coba buat projek baru ini via CLI atau Folder baru pada project.
```bash
mkdir hello
cd hello
go mod init hello
```
Bisa dilihat pada command di atas ada direktori `hello`, dibuat. Setelah masuk ke direktori tersebut, perintah `go mod init project-pertama` dijalankan. Dengan ini maka kita telah menginisialisasi direktori `hello` sebagai sebuah project Go dengan nama `hello` (kebetulan di sini nama projek sama dengan nama direktori-nya).

Skema penulisan command `go mod`:
```bash
go mod init <nama-project>
go mod init hello
```

Untuk nama project, umumnya adalah disamakan dengan nama direktori, tapi bisa saja sebenarnya menggunakan nama yang lain.

> Nama project dan Nama module merupakan istilah yang sama.

Eksekusi perintah `go mod init` menghasilkan satu buah file baru bernama `go.mod`. File ini digunakan oleh `Go toolchain` untuk menandai bahwa folder di mana file tersebut berada adalah folder projek. Jadi jangan di hapus ya file tersebut.

Sebenarnya selain `Go Modules`, setup projek di Go juga bisa menggunakan `$GOPATH`. Tapi inisialisasi project dengan `GOPATH` sudah **outdate** dan kurang dianjurkan untuk projek-projek yang dikembangkan menggunakan Go versi terbaru (1.14 ke atas)

## Membuat File Baru
Buatlah file baru dengan nama `hello.go` dan isi program seperti dibawah ini.
```go
package hello

func Hello() string {
	return "Hello, world."
}
```
Sekarang, direktori berisi sebuah paket, bukan sebuah modul, karena tidak ada berkas go.mod. Jika direktori sekarang yaitu /home/gopher/hello dan menjalankan go test, maka akan muncul:

```go
➜  hello git:(main) ✗ go test
PASS
ok      github.com/hello        0.493s
```

Berkas `"go.mod"` hanya muncul di akar direktori dari modul. Paket-paket di dalam sub-direktori memiliki path impor yang terdiri dari path modul ditambah dengan path ke sub direktori. Sebagai contohnya, jika kita buat sub direktori "world", kita tidak perlu menjalankan `"go mod init"` di dalamnya. Paket tersebut akan secara otomatis dikenal sebagai bagian dari modul "github.com/hello", dengan path impor "github.com/hello/world"

## Menambahkan Dependensi

Motivasi utama dari Go modul adalah untuk meningkatkan pengalaman dari menggunakan (yaitu, menambahkan sebuah dependensi) kode yang ditulis oleh pengembang lainnya.

Mari kita coba perbarui `"hello.go"` supaya mengimpor `"rsc.io/quote"` dan menggunakannya untuk mengimplementasikan fungsi Hello:

```go
package hello

import "rsc.io/quote"

func Hello() string {
	return quote.Hello()
}
```
Sekarang mari kita jalankan tes kembali:
```bash
➜  hello git:(main) ✗ go mod vendor
go: finding module for package rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: found rsc.io/quote in rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
```

Perintah go menangani impor dengan menggunakan versi dependensi modul tertentu yang didaftarkan dalam go.mod. Saat ia menemui sebuah import dari sebuah paket yang tidak ditemukan dalam go.mod, perintah go otomatis mencari modul yang berisi paket tersebut dan menambahkannya ke go.mod, menggunakan versi yang terakhir. ("Terakhir" didefinisikan sebagai versi terakhir yang di tag sebagai stabil — yang bukan pra-rilis, atau versi pra-release terakhir yang di tag, atau versi terakhir yang tidak di tag.) Dalam contoh di atas, "go test" menangani impor yang baru "rsc.io/quote" ke modul "rsc.io/quote v1.5.2". Ia juga mengunduh dua dependensi yang digunakan oleh "rsc.io/quote", yaitu "rsc.io/sampler" dan "golang.org/x/text". Hanya dependensi langsung saja yang dicatat dalam berkas "go.mod":
```bash
➜  hello git:(main) ✗ cat go.mod 
module example.com/hello

go 1.17

require rsc.io/quote v1.5.2

require (
        golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
        rsc.io/sampler v1.3.0 // indirect
)
```

Ingatlah bahwa walaupun perintah go membuat penambahan dependensi baru dengan cepat dan mudah, ia ada "biaya"-nya. Modul anda sekarang benar-benar bergantung pada dependensi baru dalam ruang yang rawan, beberapa hal harus diperhatikan seperti ketepatan, keamanan, dan lisensi. Untuk pertimbangan lebih lanjut, lihat artikel dari Russ Cox, Our Software Dependency Problem.

Selain `"go.mod"`, perintah go juga membuat sebuah berkas bernama "go.sum" yang berisi hash kriptografi dari isi modul pada versi tertentu:
```bash
➜  hello git:(main) ✗ cat go.sum 
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=
```

## Memperbarui dependensi

Dengan Go modul, versi-versi diacu dengan tag versi semantik. Sebuah versi semantik memiliki tiga bagian utama: mayor, minor, dan patch (tambalan). Misalnya, untuk v0.1.2, versi mayor adalah 0, versi minor adalah 1, dan versi tambalan adalah 2. Mari kita lihat bagaimana memperbarui beberapa versi minor. Dalam seksi selanjutnya, kita akan melihat bagaimana melakukan pembaruan versi mayor.

```bash
➜  hello git:(main) ✗ go get golang.org/x/text 
go get: upgraded golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c => v0.3.7
```

Sekarang mari kita coba perbarui `"rsc.io/sampler"` ke versi minor. Dimulai dengan cara yang sama, dengan menjalankan "go get" dan menjalankan tes:
```bash
➜  hello git:(main) ✗ go get rsc.io/sampler
go: downloading rsc.io/sampler v1.99.99
go get: upgraded rsc.io/sampler v1.3.0 => v1.99.99
```
dan kita coba lagi untuk `go test`
```bash
➜  hello git:(main) ✗ go test      
--- FAIL: TestHello (0.00s)
    hello_test.go:8: Hello() = "99 bottles of beer on the wall, 99 bottles of beer, ...", want "Hello, world."
FAIL
exit status 1
FAIL    example.com/hello       0.386s
```

Oo! Ternyata tes gagal, memperlihatkan bahwa versi terakhir dari "rsc.io/sampler" tidak kompatibel dengan kebutuhan kita.

Kita telah menggunakan v1.3.0; v1.99.99 jelas tidak bisa digunakan. Mungkin kita bisa mencoba v1.3.1:
```bash
➜  hello git:(main) ✗ go get rsc.io/sampler@v1.3.1
go: downloading rsc.io/sampler v1.3.1
go get: downgraded rsc.io/sampler v1.99.99 => v1.3.1
```
Kita coba ulang lagi untuk `go test`, maka hasilnya seperti dibawah ini
```bash
➜  hello git:(main) ✗ go test
PASS
ok      example.com/hello       0.130s
```
> Perhatikan perintah eksplisit `"@v1.3.1"` pada argumen `"go get"`. Pada umumnya, setiap argumen yang dikirim ke `"go get"` dapat menerima versi eksplisit; jika kosong maka dianggap sebagai `"@latest"`, yang berarti akan diubah ke versi terakhir seperti yang telah dijelaskan sebelumnya.

## Menambahkan Sebuah Dependensi dengan versi Mayor
Mari kita tambahkan satu fungsi pada program yang sudah kita buat yaitu di folder `hello` sebagai berikut:
```go
package hello

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	return quote.Hello()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
```
dalam hal ini kita menambahkan fungsi dari dependency yang sama tetapi menggunakan versi 3. Pada versi 3 ini, kita menggunakan library `quote` yang berhubungan dengan cuncurrency Go. Nanti pada sesi selanjutnya akan kita bahas.

Kemudian kita tambah sebuah berkas pengujian "hello_test.go":

```go
func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
```
setelah itu kita jalankan kembali perintah ini.
```bash
go test
PASS
ok      example.com/hello       0.890s
```
Maka skrng kita telah menggunakan 1 dependency dengan 2 versi yang berbeda yaitu `rsc.io/quote` dan `rsc.io/quote/v3`.

Pada go kita memperbolehkan pada saat `build` menggunakan beberapa versi mayor. Misalkan pada program yang telah kita buat, itu memiliki 2 versi yang berjalan bersamaan dalam satu projek. Aturan dalam go versi ini biasanya menggunakan path yang berbeda agar jika dependency ini saling keterkaitan dengan module lain bisa lebih mudah untuk melakukan migrasi-nya.

## Memperbarui dependensi ke versi mayor
Setelah kita menggunakan fungsi yang ada di versi mayor terbaru, maka saatnya fungsi yang lain juga kita harus melakukan migrasi atau `update` agar lebih mengikuti dependecy yang terbaru. Biasanya dependency yang terbaru memiliki informasi yang lengkap, perbaikan `bug` atau restruktur fungsi agar lebih bisa digunakan untuk fungsi umum dalam hal ini `general`.

Baiklah, pada program sebelumnya kita akan merubah dependency dari `v1.5.3` ke dalam versi mayor 3. Berikut dependency yang perlu kita perbaharui.
```go
package hello

import (
	 "rsc.io/quote/v3"
)

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
```

Dengan perubahan ini, maka tidak perlu lagi memberi nama pada impor dan memiliki 2 import. 

## Menghapus dependensi yang sudah tidak digunakan

Jika program telah berjalan dengan baik dan tidak ada masalah saat dijalankan, maka kita perlu juga untuk membersihkan beberapa dependency yang tidak terpakai dengan cara perinah `go mod tidy` dan `go mod vendor` untuk memastikan dependency yang terdaftar sudah sesuai.

```bash
➜  hello git:(main) ✗ cat go.mod 
module example.com/hello

go 1.17

require rsc.io/quote/v3 v3.1.0

require (
        golang.org/x/text v0.3.7 // indirect
        rsc.io/sampler v1.3.1 // indirect
)
```

Maka, bisa kita lihat dependency untuk `quote` versi terlama sudah terhapus karena sudah tidak terpakai lagi dalam program ini.

### Catatan
Perintah untuk menginisialisasi program menggunakan gomodule
```bash
go mod init <nama-proejek>
```

Perintah melakukan `build` program sekaligus untuk melakukan `test`.
```bash
go build 
go test
```

Perintah untuk memperbarui dependency
```bash
go get <path-package-dependency>
```

Perintah untuk menghapus dependency yang tidak terpakai
```bash
go mod tidy
```
