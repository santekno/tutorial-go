# Channel Select

Adanya channel sangat membantu kita untuk mengelola goroutine yang sedang berjalan pada program kita. Ada kalanya kita juga perlu mengelola goroutine banyak dan dibutuhkan banyak channel juga. Maka, disinilah kegunaan `select`. `select` memudahkan kita mengontrol komunikasi data lewat channel. Cara penggunannya sama seperti seleksi kondisi `switch`.


Program pencarian **rata-rata** dan **nilai tertinggi** berikut merupakan contoh sederhana penerapan select dalam channel. Akan ada 2 buah goroutine yang masing-masing di-handle oleh sebuah channel. Setiap kali goroutine selesai dieksekusi, akan dikirimkan datanya ke channel yang bersangkutan. Lalu dengan menggunakan select, akan dikontrol penerimaan datanya.
Pertama, kita siapkan terlebih dahulu 2 fungsi yang akan dieksekusi sebagai goroutine baru. Fungsi pertama digunakan untuk mencari **rata-rata**, dan fungsi kedua untuk penentuan **nilai tertinggi** dari sebuah slice.


Baiklah, kita akan mencoba mengimplementasikan program sederhana untuk pencarian **rata-rata** dan **nilai tertinggi** dari suatu bilangan array. Pada program ini akan kita definisikan 2 buah goroutine yang masing-masing memiliki fungsi yang akan di-handle oleh sebuah channel. 

Setiap kali goroutine selesai dieksekusi maka akan dikirimkan datanya ke channel yang bersangkutan, lalu dengan menggunakan select akan dikontrol penerimaan datanya.

Kita siapkan terlebih dahulu 2 fungsi yang akan dieksekusi sebagai goroutine baru yaitu fungsi pertama mencari **rata-rata** dan fungsi kedua menentukan **nilai tertinggi** dari sebuah `array slice`.
```go
package main

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}
```
Kedua fungsi diatas akan kita eksekusi pada fungsi `main` sebagai goroutine baru. Setelah itu kedua fungsi tersebut akan mengirimkan datanya ke dalam channel yang sudah ditentukan, yang mana pada program ini kita akan membedakan variabel penampungnya. `ch1` menampung untuk nilai rata-rata, `ch2` untuk hasil data nilai tertinggi.

Setelahnya, mari kita buat fungsi `main` dibawah ini.
```go
func main() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers : ", numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}
}
```

Pada kode di atas, transaksi pengiriman data pada channel ch1 dan ch2 dikontrol menggunakan select . Terdapat 2 buah case kondisi penerimaan data dari kedua channel tersebut.
Kondisi case avg := <-ch1 akan terpenuhi ketika ada penerimaan data dari channel ch1 , yang kemudian akan ditampung oleh variabel avg .
Kondisi case max := <-ch2 akan terpenuhi ketika ada penerimaan data dari channel ch2 , yang kemudian akan ditampung oleh variabel max .
Karena ada 2 buah channel, maka perlu disiapkan perulangan 2 kali sebelum penggunaan keyword select .

Pada kode di atas, pengiriman data pada channel `ch1` dan `ch2` akan dikontrol oleh `select`. Terdapat 2 case kondisi dimana penerimaan data dari kedua tersebut akan dikirim melalui channel tersebut.
* Kondisi `case avg := <-ch1` akan terpenuhi ketika penerimaan data dari channel `ch1` yang kemudian akan ditampung pada variabel `avg`.
* Kondisi `case max := <-ch2` akan terpenuhi ketika penerimaan data dari channel `ch2` yang kemudian akan ditampung pada variabel `max`.

Dikarenakan ada 2 buah channel maka kita perlu disiapkan untuk melakukan perulangan sebanyak 2 kali sebelum penggunaan keyword `select`.
```bash
➜  channel-select git:(main) ✗ go run main.go
numbers : [3 4 3 5 6 3 2 2 6 3 4 6 3]
Max     : 6 
Avg     : 3.85 
➜  channel-select git:(main) ✗ go run main.go
numbers : [3 4 3 5 6 3 2 2 6 3 4 6 3]
Avg     : 3.85
Max     : 6 
```

## Channel - Range and Close
Penerimaan data lewat channel yang dipakai oleh banyak goroutine, akan lebih mudah dengan memanfaatkan keyword `for - range`. Hal ini bisa kita terapkan untuk melakukan perulangan tanpa henti.

Perulangan tersebut tetap berjalan meski tidak ada transaksi pada channel dan hanya akan berhenti jika status channel berubah menjadi closed atau sudah ditutup. Fungsi `close` digunakan utuk menutup channel.

Channel yang sudah ditutup tidak akan bisa digunakan kembali untuk menerima dan mengirim data sehingga menjadikan `for - range` pun akan ikut berhenti juga.

Misalkan kita akan membuat contoh implementasinya seperti dibawah ini.
Kita akan membuat suatu fungsi `sendMessage` dengan parameter channel yang mana jika fungsi itu dijalankan akan melakukan perulangan sebanyak 20 kali dan setelah semua terkirim maka akan diakhiri dengan channel `close`.
```go
func sendMessage(ch chan<- string) {
  for i := 0; i < 20; i++ {
      ch <- fmt.Sprintf("data %d", i)
  }
  close(ch) 
}
```
Siapkan juga fungsi `printMessage()` untuk handle penerimaan data. Didalamnya, channel akan di-looping menggunakan `for - range` yang kemudian ditampilkan datanya.
```go
func printMessage(ch <-chan string) {
    for message := range ch {
        fmt.Println(message)
    }
}
```

Setelah itu kita akan buat channel baru di fungsi main dengan menjalankan `sendMessage()` sebagai goroutine dan fungsi `printMessage` juga akan kita jalankan. Maka, dengan ini kita akan mengirimkan 20 data lewat goroutine baru dan juga akan diterima oleh goroutine juga.
```go
func main() {
    runtime.GOMAXPROCS(2)
    var messages = make(chan string)
    go sendMessage(messages)
    printMessage(messages)
}
```

Jika 20 data itu sukses dikirim dan diterima maka channel `ch` akan dimatikan `close`, sehingga perulangan data channel dalam `printMessage()` juga akan berhenti.
```bash
➜  channel-range-close git:(main) ✗ go run main.go 
data 0
data 1
data 2
data 3
data 4
data 5
data 6
data 7
data 8
data 9
data 10
data 11
data 12
data 13
data 14
data 15
data 16
data 17
data 18
data 19
```

## Channel Direction
Golang memiliki keunikan dalam fitur parameter channel yang sudah disediakan.
Level akses channel bisa ditentukan, apakah kita hanya sebagai penerima, pengirim atau bahkan sekaligus penerima dan pengirim. Konsep ini disebut dengan **channel direction**.

Cara pemberian level akses ini dengan menambahkan tanda `<-` sebelum atau setelah keyword `chan`. Untuk lebih jelasnya bisa dilihat di tabel berikut.

| Sintaks | Penjelasan |
| :- | :--------- |
| `ch chan string` | Parameter `ch` bisa digunakan untuk mengirim dan menerima data |
| `ch chan<- string` | Parameter `ch` hanya bisa digunakan untuk mengirim data |
| `ch <-chan string` | Parameter `ch` hanya bisa digunakan untuk menerima data |

Secara default channel akan memiliki kemampuan untuk mengirim dan menerima data. Agar channel tersebut hanya bisa mengirim atau menerima saja, maka kita perlu memanfaatkan simbol `<-`.

Sebagai contoh fungsi `sendMessage(ch chan<- string)` yang parameter `ch` dideklarasikan dengan level akses untuk pengiriman data saja. Channel tersebut hanya bisa digunakan untuk mengirim data, contohnya: `ch <- fmt.Sprintf("data %d", i)`.

Dan sebaliknya pada fungsi `printMessage(ch <-chan string)`, channel `ch` hanya bisa digunakan untuk menerima data saja.

## Channel - Timeout
Mendefinisikan channel agar bisa mengontrol penerimaan data dari channel berdasarkan waktu diterimanya maka kita perlu pengecekan durasi timeout dengan kita bisa menentukan sendiri. Misalkan ketika tidak ada aktivitas penerimaan data selama durasi tertentu, maka akan memicu `callback` yang isinya juga ditentukan sendiri.

Berikut adalah program sederhana tentang pengaplikasian `timeout` pada channel. Sebuah goroutine baru dijalankan dengan tugas mengirimkan data setiap interval tertentu, dengan durasi interval-nya adalah acak/random.
```go
package main

import (
	"math/rand"
	"time"
)

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second) // fungsi tertentu yg proses-nya lama, kadang2
	}
}
```
Selanjutnya, disiapkan perulangan tanpa henti, pada setiap perulangannya ada seleksi kondisi channel menggunakan `select`.
```go
func retreiveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout. no activities under 5 seconds")
			break loop
		}
	}
}
```
Pada fungsi `retrieveData` terdapat 2 kondisi yang mana ini terjadi jika
* `case data := <- messages:` maka akan terpenuhi ketika ada serah terima data pada channel `messages`
* `case <- time.After(time.Second * 5):` maka akan terpenuni ketika tidak ada aktivitas penerimaan data dari channel dalam durasi waktu 5 detik.

Maka diakhir kita akan eksekusi fungsi tersebut pada fungsi `main`. Beriku ini lebih lengkapnya.
```go
func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var messages = make(chan int)
	
	go sendData(messages)
	retreiveData(messages)
}
```
Akan muncul output setiap kali ada penerimaan data dengan delay waktu acak. Ketika tidak ada aktifitas pada channel dalam durasi 5 detik, perulangan pengecekkan channel akan dihentikan.
```bash
➜  channel-timeout git:(main) ✗ go run main.go
receive data "0"
receive data "1"
timeout. no activities under 5 seconds
➜  channel-timeout git:(main) ✗ go run main.go
receive data "0"
timeout. no activities under 5 seconds
➜  channel-timeout git:(main) ✗ go run main.go
receive data "0"
timeout. no activities under 5 seconds
➜  channel-timeout git:(main) ✗ go run main.go
receive data "0"
receive data "1"
receive data "2"
timeout. no activities under 5 seconds
```