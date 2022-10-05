# Hashes & Cryptography

Fungsi `hash` mengambil satu set data dan menguranginya ke ukuran tetap yang lebih kecil. **Hash** sering digunakan dalam pemrograman untuk segala hal mulai dari mencari data hingga mendeteksi perubahan dengan mudah. Fungsi `hash` di Go dibagi menjadi dua kategori yaitu `kriptografi` dan `non-kriptografi`. Fungsi `hash` `non-kriptografi` dapat ditemukan di bawah paket `hash` dan termasuk `adler32`, `crc32`, `crc64` dan `fnv`. 

Berikut ini contoh menggunakan `crc32`:
```go
package main
import (
     "fmt"
     "hash/crc32"
)
func main() {
     h := crc32.NewIEEE()
     h.Write([]byte("test"))
     v := h.Sum32()
     fmt.Println(v)
}
```

**Hash** `crc32` mengimplementasikan interface `Writer`, sehingga kita dapat menulis `byte` padanya seperti `Writer` lainnya. Setelah kita menulis semua yang kita inginkan, kita memanggil `Sum32()` untuk mengembalikan `uint32`. Penggunaan umum untuk `crc32` adalah untuk membandingkan dua file. Jika nilai `Sum32` untuk kedua file sama, kemungkinan besar (meskipun tidak 100% yakin) bahwa file tersebut sama. Jika nilainya berbeda maka file tersebut pasti tidak sama.

Berikut membandingkan dua file dengan menggunakan `crc32` dibawah ini.
```go
package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32, error) {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(bs)
	return h.Sum32(), nil
}
func main() {
	h1, err := getHash("test1.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}
	fmt.Println(h1, h2, h1 == h2)
}
```
Kita coba skenariokan misalkan kita isi file `test1.txt` dan `test2.txt` sebagai berikut:
- `test1.txt` diisi `hello world` dan `test2.txt` diisi dengan `hello world santekno` --> ekspektasi data dari kedua file tersebut berbeda maka return (false)
- `test1.txt` diisi `hello world` dan `test2.txt` diisi dengan `hello world` --> ekspektasi data dari kedua file tersebut sama, maka return (true)

## Cryptography