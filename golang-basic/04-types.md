# Tipe dalam Golang

# Numbers
## Integers 
Terdapat tipe integer yang bisa kita gunakan diantaranya yaitu `uint8`, `uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32` dan `int64`. Berikut ini contoh penggunaan variabel tipe integer pada program go:
```go
package main

import "fmt"
func main() {
    fmt.Println("1 + 1 =", 1 + 1)
}
```
Jika kamu menjalankan [program ini](golang-basic/04-types/integer.go) dan kamu bisa melihat outputnya seperti dibawah ini:
```bash
$ go run main.go 
1+1=2
```

Dari program tersebut kita akan mencetak string berupa `1 + 1=` dan operasi `1 + 1` pada aritmatika yang akan mengoperasikan penjumlahan yaitu dengan hasil `2`. 

## Floating 
Terdapat tipe float yang bisa kita gunakan diantaranya yaitu `float32`, `float64`,`complex64` dan `complex128`.

Bagaimana jika kita ubah operasinya itu menjadi `1.0 + 1.0`? Apakah hasil operasi aritmatika tersebut akan sama? 

Tentu akan berbeda karena `go` mendefinisikan `.0` itu sebagai tipe data `float`.

Selain itu kita juga perlu tahu bahwa dalam go memiliki beberapa operasi dasar aritmatika sebagai berikut:
| x | Operation |
| :-: | :--------- |
| + | penjumlahan |
| - | pengurangan |
| * | perkalian |
| / | pembagian |
| % | modulo |

## Strings
Jika kita akan melakukan cetak `string` ada 2 cara yaitu dengan menggunakan tanda petik satu dan tanda petik dua. Misalkan seperti ini ``Hello World`` dan `"Hello World"` bisa digunakan. Untuk menambahkan `new line` pada string yang akan dicetak kita bisa tambahkan `\n` dan untuk `tab character` bisa menambahkan `\t`.

Kita coba modifikasi program `hello-world.go` menjadi program seperti dibawah ini:
```go
package main

import "fmt"

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
}
```
Beberapa hal yang perlu diperhatikan:
* Spasi pada string akan dihitung sebagai 1 karakter sehingga pada baris 1 akan menghasilkan `11` yang mana fungsi `len` digunakan untuk mengetahui jumlah karakter dalam string
* String memulai index dari 0 bukan dari 1. `[1]` mencari element ke-2 dari `Hello World` yaitu karakter `e`. Perhatikan bahwa yang dicetak oleh program baris 2 bukan `e` tetapi `101` saat menjalankan program ini. Ingat! program akan mencetak karakter yang diwakili dengan byte dan `byte` adalah bilangan bulat. 
(ASCII dari huruf `e` adalah `101`)
* Penggabungan menggunakan simbol yang sama dengan penambahan. Kompiler Go mencari tahu apa yang harus dilakukan berdasarkan jenis argumen. Karena `+` adalah string, kompiler menganggap yang kita maksud adalah penggabungan dan bukan penambahan.


## Boolean
Boolean mewakili bit yang mana memiliki nilai 1 bit integer yang mana merepresentasikan sebagai `true` dan `false`. Ada 3 operasi logika yang bisa kita gunakan yaitu:
| x | Operation |
| :-: | :--------- |
| && | and |
| `||` | or |
| ! | not |

Berikut ini contoh [program ini]() agar kamu bisa lebih memahami.
```go
package main

import "fmt"

func main() {
    fmt.Println(true && true)
    fmt.Println(true && false)
    fmt.Println(true || true)
    fmt.Println(true || false)
    fmt.Println(!true)
}
```
setelah menjalankan program ini maka kita akan melihat output seperti ini
```bash
$ go run boolen.go
true
false
true
true 
false
```

Kita bisa lihat juga tabel logika `and` dibawah ini
| Ekspresi | Value |
| :-: | :--------- |
| true && true | true |
| true && false | false |
| false && true | false |
| false && false | false |

Tabel logika `or`
| Ekspresi | Value |
| :-: | :--------- |
| true \|\| true | true |
| true \|\| false | true |
| false \|\| true | true |
| false \|\| false | false |

Tabel Logika `not`
| Ekspresi | Value |
| :-: | :--------- |
| !true | false |
| !false | true |