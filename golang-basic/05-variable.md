# Variabel & Constants

## Variable
Berikut contoh implementasi dasar penggunaan variable pada program dibawah ini
```go
package main

import "fmt"

func main() {
	var x string
	x = "Hello World"
	fmt.Println(x)
}

```
Dalam program ini kita akan mencetak string `Hello World` dengan menyimpan terlebih dahulu string ke dalam suatu variabel yaitu `x` dan selanjutnya mencetak variabel tersebut menggunakan fungsi `fmt.Println()`.

Jika kita mau menggabungkan variable string, kita bisa menggunakan operation `+` untuk menggabungkannya. Misalkan contoh program ada dibawah ini:
```go
package main

import "fmt"

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
}

```
setelah program dijalankan maka outputnya seperti dibawah ini
```
$ go run main.go 
first 
first second
```

Kita juga bisa menggunakan simbol persamaan yang berbeda `==`. (Dua tanda sama dengan di sebelah satu sama lain) `==` adalah operator seperti + dan mengembalikan boolean. Sebagai contoh:
```go
package main

import "fmt"

func main() {
  var x string = "hello"
  var y string = "world"
  fmt.Println(x == y)
}
```
Program ini akan mengeluarkan output `false`.
Sedangkan jika kita ganti variabel string `y` menjadi `hello` seperti ini
```go
  ...
  var y string = "world"
  ...
```
maka program akan mengeluarkan output `true`.

Didalam Go juga kita bisa membuat variabel dengan lebih singkat dan mudah sebagai contoh:
```go
  x := "Hello World"
```
Tanda petik `:` sebelum sama dengan `=` menandakan bahwa kita akan menyimpan isi ke dalam variabel `x` dan tipe dari variabel `x` pun akan disesuaikan oleh compiler yaitu berupa tipe `string`.

Misalkan kita buat lagi variabel seperti dibawah ini
```go
  x := 5
  fmt.Println(x)
```
Maka compiler pun akan menganggap variabel `x` ini berupa tipe `integer`

## Bagaimana mendefinisikan suatu variabel
Memberi nama variabel dengan benar merupakan bagian penting dari pengembangan perangkat lunak. Nama harus dimulai dengan huruf dan dapat berisi huruf, angka, atau simbol _ (garis bawah). Kompilator Go tidak peduli apa yang Anda beri nama variabel sehingga nama tersebut dimaksudkan untuk keuntungan Anda (dan lainnya). Pilih nama yang secara jelas menggambarkan tujuan variabel. Misalkan kita memiliki yang berikut:
```go
  x := "Kitty"
  fmt.Println("Nama kucing saya adalah ", x)
```
Dalam hal ini `x` bukanlah nama yang sangat baik untuk sebuah variabel. Nama yang lebih baik adalah:
```go
  name := "Kitty"
  fmt.Println("Nama kucing saya adalah ", name)
```
atau bisa juga
```go
  catName := "Kitty"
  fmt.Println("Nama kucing peliharaanku adalah ", catName)
```

## Scope
Kembali lagi ke program misalkan kita mendefinisikan variabel seperti dibawah ini
```go
package main

import "fmt"

func main() {
  var x string
  x = "Hello World"
	fmt.Println(x)
}
```
lalu kita edit menjadi seperti ini
```go
package main

import "fmt"

var x string = "Hello World"

func main() {
	fmt.Println(x)
}
```
Perhatikan bahwa kita memindahkan variabel di luar fungsi utama. Ini berarti bahwa fungsi lain dapat mengakses variabel ini:
```go
package main

import "fmt"

var x string = "Hello World"

func main() {
	fmt.Println(x)
}

func f(){
	fmt.Print(x)
}
```
Sekarang fungsi `f()` akan mengakses variabel `x` ini jika fungsi tersebut dipanggil. Lalu kita coba lihat program ini seperti ini:
```go
package main

import "fmt"


func main() {
  var x string = "Hello World"
	fmt.Println(x)
}

func f(){
	fmt.Print(x)
}
```
Jika kamu akan menjalankan ini maka program tersebut akan error.
```bash
  .\main.go:11: undefined: x
```

Kompiler memberi tahu bahwa variabel `x` di dalam fungsi `f()` tidak ada. Itu hanya ada di dalam fungsi utama. Rentang tempat di mana diizinkan untuk menggunakan `x` disebut **ruang lingkup variabel**.

Pada dasarnya, variabel ada di dalam kurung kurawal terdekat `{}` (blok) termasuk kurung kurawal bersarang (blok), tetapi tidak di luarnya. Lingkup dapat sedikit membingungkan pada awalnya; seperti yang kita lihat lebih banyak contoh Go, itu akan menjadi lebih jelas.

## Constants
Go juga memiliki dukungan untuk *konstanta*. Konstanta pada dasarnya adalah variabel yang nilainya tidak dapat diubah nanti. Mereka dibuat dengan cara yang sama seperti membuat variabel tetapi alih-alih menggunakan kata kunci `var`, kami menggunakan kata kunci `const`.
```go
package main

import "fmt"

func main() {
	const x string = "Hello World"
	fmt.Println(x)
}
```
Lalu kita coba menambahkan 1 baris dibawah constant seperti ini
```go
  const x string = "Hello World"
  x = "Some other string"
```
maka compiler akan mengeluarkan error seperti ini
```bash
.\main.go:7: cannot assign to x
```

Penggunaan `constant` hanya digunakan untuk variable yang tidak akan di `write` setiap kali. Misalnya `Pi` dalam matematika ini bisa dijadikan tipenya sebagai constant.

## Multiple Definition
Go juga bisa mempermudah penginisialisasian variabel dengan cara multiple definition seperti contoh dibawah ini.
```go
var ( 
  a = 5
  b = 10
  c = 15 
)
```
Multiple definision ini bisa digunakan untuk `var` dan `const`.

## Program untuk membaca variable
Berikut ini contoh menginput variabel dengan scan program
```go
package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}

```

maka output:
```bash
Enter a number: 4
8
```