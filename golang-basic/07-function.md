# Function

## Second Function
Terdiri dari `input`, `output`.

Sampai sekarang program Go itu kita bisa gunakan hanya pada fungsi main saja seperti ini.
```go
  func main() {}
```
Misalkan pada program sebelumnya kita mengoperasikan penjumlahan, dan rata-rata dalam suatu array, nah agar lebih modular lagi kita akan ganti menjadi fungsi terpisah seperti ini.
```go
func average(xs []float64) float64 {
     total := 0.0
     for _, v := range xs {
           total += v
}
     return total / float64(len(xs))
}
```
Maka kita akan coba-coba bagaimana suatu variable bisa bekerja dalam fungsi-fungsi dan dependency yang terkait.
Program A
```go
func f() {
	fmt.Println(x)
}
func main() {
	x := 5
	f()
}
```
Program B
```go
func f(x int) {
     fmt.Println(x)
}
func main() {
x := 5
f(x) }
```

Program C
```go
var x int = 5
func f() {
     fmt.Println(x)
}
func main() {
     f()
}
```
Program D
```go
func main() {
	fmt.Println(f1())
}
func f1() int {
	return f2()
}
func f2() int {
	return 1
}
```

## Multiple Returning
Jika dalam pemrograman yang lain hanya bisa membalikan hasil cuma satu data tetapi dalam Go, memiliki keunikan tersendiri yaitu kita bisa melakukan output `return` yang lebih dari satu. Agar lebih bisa dipahami lihat dibawah ini.
```go
package main

import "fmt"

func f() (int, int) {
	return 5, 6
}
func main() {
	x, y := f()
	fmt.Println(x, y)
}
```

## Variadic Function
Yang paling spesial juga dalam pemrograman Go yaitu support untuk Variadic, yaitu multiple parameter.
```go
package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func main() {
	fmt.Println(add(1, 2, 3))
}
```

## Closure
Biasanya sering kita lihat pada pemrograman `javascript` atau `typescript`. Nah di Golang juga memiliki hal yang canggih seperti Closure ini. Berikut contoh programnya dibawah ini.
```go
package function

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
}
```
Mengirim variable ke function closure juga bisa kita gunakan misalkan seperti ini
```go
package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
```

Contoh program lain diluar function bisa kita buat seperti ini
```go
package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}
```
## Recursion
Memanggil fungsinya sendiri juga dalam Go bisa kita lakukan, misalkan contoh program untuk menghitung faktorial.
```go
package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(uint(3)))
}
```

## Defer, Panic, Recover
defer adalah suatu pemanggilan yang nantinya akan diakhirkan. contoh bisa dilihat dibawah ini
```go
package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {
	defer second()
	first()
}
```
Lalu untuk mendefinisikan suatu program bisa berhenti dan bisa recover lagi dalam Go juga bisa kita pakai. Contohnya seperti dibawah ini.

```go
package main
import "fmt"
func main() {
     defer func() {
           str := recover()
           fmt.Println(str)
     }()
     panic("PANIC")
}
```