# Library Input / Output

Sebelum mempelajari libary IO Golang, banyak sekali fungsi terdapat dalam library IO tetapi yang paling utama ada 2 interface fungsi yaitu `Reader` dan `Writer`. `Reader` biasanya digunakan untuk membaca data dari suatu file atau beberapa input/output yang disediakan. Sedangkan `Writer` adalah fungsi yang nantinya akan menulis data ke suatu file atau input/output yang disediakan oleh sistem kita. 

## Interface `Reader`
Interface Reader ini banyak sekali tipe implementer fungsinya tetapi yang akan kita coba adalah reader yg diambil dari file. Kita bisa lihat contoh kodenya dibawah ini.
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal("program broke")
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("program read broke")
	}

	fmt.Println(bs)
	fmt.Println(string(bs))
}
```

Dan isi dari file `file.txt` sebagai berikut
```txt
Hello World
```

Setelah kita coba lihat code diatas, kebayang ada `os.Open` adalah fungsi yang digunakan untuk membuka dan mencari file sehingga nanti data tersebut akan kita baca dengan menggunakan fungsi `ioutil.ReadAll(f)`.

Jika kita lihat code diatas, pengembalian dari fungsi `ioutil.ReadAll` ini ditangkap oleh variable `bs` yang mana kita akan cetak dan menampilkan isinya. Secara default variable `bs` ini tipe datanya `byte` sehingga kita perlu menambahkan fungsi `string(bs)` agar bisa terbaca.

## Interface `Writer`
Interface `Writer` digunakan untuk kita bisa menulis ke dalam satu file atau beberapa file sekaligus. Kita akan coba lihat code dibawah ini contoh dasar penggunaan interface `Writer`

```go
package main

import (
	"log"
	"os"
)

func main() {
	dst, err := os.Create("create.txt")
	if err != nil {
		log.Fatal("program broke")
	}
	defer dst.Close()

	bs := []byte("hello world, Santekno")

	_, err = dst.Write(bs)
	if err != nil {
		log.Fatal("program broke")
	}
}
```
Jika kita jalankan program diatas ini maka akan terbuat file `create.txt` dengan isi dari file tersebut berupa string `hello world, Santekno`. Program ini sederhana dimana kita bisa tahu file yang di create dengan nama apa dan isi dari file tersebut bisa kita kontrol apa saja.