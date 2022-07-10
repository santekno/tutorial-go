# Channel

Channel adalah penghubung goroutine yang dari satu ke yang lainnya. Channel ini sifat-nya `synchronous` karena ada blocking.

Channel bisa didefinisikan dengan bentuk variabel dengan keyword `chan`. Variabel ini memiliki tugas untuk mengirim dan menerima data.

Berikut contoh implementasi channel yang memiliki 3 goroutine baru dieksekusi dan masing-masing gouroutine melakukan proses data lewat channel. Data tersebut akan diterima 3 kali di goroutine utama (`main`).
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("captain america")
	go sayHelloTo("captain marvel")
	go sayHelloTo("black panter")
	
  var message1 = <-messages
	fmt.Println(message1)
	
  var message2 = <-messages
	fmt.Println(message2)
	
  var message3 = <-messages
	fmt.Println(message3)
}
```

Pada kode di atas, variabel `messages` dideklarasikan bertipe `channel string`. Cara inisialisasi variabel `channel` yaitu dengan menuliskan `make` dengan keyword `chan` diikuti dengan tipe data channel yang diinginkan.
```go
var messages = make(chan string)
```
Pada program ini, kita buat juga fungsi `closure sayHelloTo` yang mengembalikan data string. Data tersebut kemudian dikirim lewat channel `messages`. Tanda `<-` jika dituliskan di sebelah kiri nama variabel maka proses tersebut sedang berlangsung proses pengiriman data dari variabel yang berada di kanan lewat channel yang berada dikiri. Pada konteks ini, variabel data dikirim lewat channel `messages`.

```go
var sayHelloTo = func(who string) {
    var data = fmt.Sprintf("hello %s", who)
    messages <- data
}
```
Fungsi `sayHelloTo` dieksekusi 3 kali sebagai `goroutine` berbeda. Menjadikan tiga proses ini berjalan secara `asynchronous` atau tidak saling tunggu.
```go
go sayHelloTo("captain america")
go sayHelloTo("captain marvel")
go sayHelloTo("black panther")
```

Dari ketiga fungsi tersebut, goroutine yang paling awal bertugas untuk mengirim data dan akan diterima datanya oleh variabel `message1` . Tanda `<-` jika dituliskan di sebelah kanan channel, menandakan proses penerimaan data dari channel yang di kanan, untuk disimpan ke variabel yang di kiri.
```go
var message1 = <-messages
fmt.Println(message1)
```

Dikarenakan `channel` bersifat `blocking` artinya statement `var message1 = <-messages` hingga setelahnya tidak akan dieksekusi sebelum ada data yang dikirim lewat channel.

Ketiga goroutine tersebut datanya akan diterima secara berurutan oleh `message1`, `message2` dan `message3` kemudian ditampilkan.
```bash
➜  channel git:(main) ✗ go run main.go 
hello black panter
hello captain marvel
hello captain america
➜  channel git:(main) ✗ go run main.go
hello black panter
hello captain marvel
hello captain america
➜  channel git:(main) ✗ go run main.go
hello black panter
hello captain marvel
hello captain america
```
Dapat dilihat bahwa output yang dikembalikan `sayHelloTo` tidak selalu berurutan, meskipun penerimaan datanya berurutan. Hal ini dikarenakan pengiriman data dari 3 goroutine yang berbeda, yang kita tidak tahu mana yang lebih dahulu dieksekusi. Goroutine yang dieksekusi lebih awal, datanya akan diterima lebih awal juga.

Sudah disinggung di atas bahwa channel ini *blocking* sehingga tidak perlu menggunakan perintah `fmt.Scanln()` untuk melakukan *blocking* tersebut.

## Channel Tipe Data Parameter

Variabel channel bisa di-passing ke fungsi lain sebagai `parameter`. Caranya dengan menambahkan keyword `chan` ketika melakukan deklarasinya.

Kita lanjutkan dengan mempraktekkan saja. Siapkan fungsi `printMessage` dengan parameter channel, lalu ambil data yang dikirimkan lewat channel tersebut untuk ditampilkan.
```go
func printMessage(what chan string) {
    fmt.Println(<-what)
}
```
Setelah itu ubah beberapa baris untuk implementasi di fungsi `main`.
```go
func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)

	for _, each := range []string{"america", "marvel", "panther"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}
	
	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}
```
Parameter `what` pada fungsi `printMessage` bertipe `channel string`. Operasi serah-terima data akan bisa dilakukan pada variabel tersebut, dan akan berdampak juga pada variabel `messages` di fungsi `main`.

Passing data bertipe channel lewat parameter secara implisit adalah `pass by reference`, yang di-passing adalah pointer-nya. Output program di atas adalah sama dengan program sebelumnya.
```go
➜  channel-parameter git:(main) ✗ go run main.go 
hello america
hello marvel
hello panther
➜  channel-parameter git:(main) ✗ go run main.go 
hello panther
hello marvel
hello america
```

Kita akan kupas tuntas program di atas bisa berjalan.

### Iterasi Data Array Pada Inisialisasi
Data `array` yang baru di-inisialisasi bisa langsung diiterasi, caranya mudah dengan menuliskanya langsung setelah keyword `range`.
```go
for _, each := range []string{"america", "marvel", "panther"} {
  // ...
}
```

## Buffered Channel
Channel secara default adalah `un-buffered`, tidak di-buffer di memori. Ketika ada goroutine yang mengirimkan data lewat channel, harus ada goroutine lain yang bertugas menerima data dari channel yang sama dengan proses serah-terima yang bersifat blocking. Maksudnya, baris kode di bagian pengiriman dan penerimaan data, tidak akan akan diproses sebelum proses serah-terima-nya selesai.

**Buffered channel** sedikit berbeda. Pada channel jenis ini, ditentukan jumlah buffer-nya. Angka tersebut akan menjadi penentu kapan kita bisa mengirimkan data. Selama jumlah data yang dikirim tidak melebihi jumlah buffer, maka pengiriman akan berjalan *asynchronous* (tidak blocking).

Ketika jumlah data yang dikirim sudah melewati batas buffer, maka pengiriman data hanya bisa dilakukan ketika salah satu data sudah diambil dari channel, sehingga ada slot channel yang kosong. Dengan proses penerimaan-nya sendiri bersifat *blocking*.

Berikut contoh penerapan **Buffered Channel** dengan membuat program bahwa pengiriman data lewat **buffered channel** adalah *asynchronous* selama jumlah data yang sedang di-buffer oleh channel tidak melebihi kapasitas buffernya.
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	messages := make(chan int, 2)
	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}
}
```
Pada kode di atas, parameter kedua fungsi `make` merepresentasikan jumlah buffere-nya. Nah, perlu kita perhatikan bahwa nilai **buffered channel** dimulai dari 0. Jadi, ketika nilainya adalah 2, berarti jumlah buffere maksimal ada 3 yaitu 0, 1, 2.

Pada contoh di atas terdapat sebuah goroutine yang berisikan proses penerimaan data dari channel `message` yang selanjutnya akan ditampilkan.

Setelah goroutine penerima data dieksekusi, maka data akan dikirimkan lewat perulangan `for`. Pada program di atas kita akan mengirim sejumlah 5 data lewat channel `message` secara sekuensial.
```bash
➜  channel-buffered git:(main) ✗ go run main.go
send data 0
send data 1
send data 2
send data 3
receive data 0
receive data 1
receive data 2
receive data 3
send data 4
```
Bisa dilihat hasilnya pada output di atas. Pengiriman data ke-4, diikuti dengan penerimaan data, dan kedua proses tersebut berjalan secara blocking.
Pengiriman data ke 0, 1, 2 dan 3 akan berjalan secara asynchronous, hal ini karena channel ditentukan nilai buffer-nya sebanyak 3 (ingat, dimulai dari 0). Pengiriman selanjutnya (ke-4 dan ke-5) hanya akan terjadi jika ada salah satu data dari 4 data yang sebelumnya telah dikirimkan, sudah diterima (dengan serah terima data yang bersifat blocking). Setelahnya, sesudah slot channel ada yang kosong, serah-terima akan kembali asynchronous.

Hasil dari output di atas kita bisa lihat bahwa pengiriman ke-4, diikuti dengan penerimaan data, dan kedua proses tersebut berjalan secara blocking. Kenapa seperti itu? Karena pada saat pengiriman data ke 0, 1, 2 dan 3 akan berjalan secara `asynchronous`, hal ini karena channal yang kita tentukan memiliki nilai buffer sebanyak 3 (dimulai dari 0).

Peingiriman selanjutnya (ke-4 dan ke-5) hanya akan terjadi jika ada salah satu data dari 4 data yang sebelumnya telah dikirimkan dan sudah diterima (serah terima data yang bersifat blocking). Sesudah slot channel kosong maka data ke-4 dan ke-5 akan di serah-terima kembali dengan `aysnchronous`.
