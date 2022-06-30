# Structs and Interfaces

Meskipun mungkin bagi kita untuk menulis program hanya menggunakan tipe data bawaan Go, pada titik tertentu itu akan menjadi sangat membosankan. Pertimbangkan program yang berinteraksi dengan bentuk:

```go
package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5
	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))
}
```

Melacak semua koordinat membuat sulit untuk melihat apa yang sedang dilakukan program dan kemungkinan akan menyebabkan kesalahan.

## Structs
Cara mudah untuk membuat program ini lebih baik adalah dengan menggunakan `struct`. Sebuah `struct` adalah tipe yang berisi bidang bernama. Misalnya kita bisa mewakili Lingkaran seperti ini:
```go
type Circle struct {
  x float64
  y float64
  r float64 
}
```
Kata kunci `type` memperkenalkan tipe baru. Diikuti dengan nama tipe (Lingkaran), kata kunci `struct` untuk menunjukkan bahwa kita mendefinisikan tipe `struct` dan daftar bidang di dalam kurung kurawal. Setiap bidang memiliki nama dan jenis. Seperti halnya fungsi, kami dapat menutup bidang yang memiliki tipe yang sama:
```go
type Circle struct {
     x, y, r float64
}
```
## Inisialiasi
Kita dapat membuat turunan dari tipe Lingkaran baru kita dalam berbagai cara:
```go
  var c Circle
```

Seperti tipe data lainnya, ini akan membuat variabel `Cirlce` lokal yang secara default disetel ke nol. Untuk `struct` nol berarti masing-masing bidang diatur ke nilai nol yang sesuai (`0` untuk int, `0,0` untuk float, `""` untuk string, `nil` untuk pointer, ...) Kita juga dapat menggunakan fungsi `new`:
```go
  c := new(Circle)
```
Ini mengalokasikan memori untuk semua bidang, menetapkan masing-masing ke nilai `nol` dan mengembalikan pointer. (`*Circle`) Lebih sering kita gunakan untuk memberi nilai pada setiap `field`. Kita bisa melakukan ini dengan dua cara. Seperti ini:
```go
  c := Circle{x: 0, y: 0, r: 5}
```
Atau kita dapat mengabaikan nama field jika kita mengetahui urutannya:
```go
  c := Circle{0, 0, 5}
```

## Fields
Kita dapat mengakses `field` menggunakan `.` operator:
```go
  fmt.Println(c.x, c.y, c.r)
  c.x = 10
  c.y = 5
```

Mari kita ubah fungsi `circleArea` sehingga menggunakan `Circle`:
```go
func circleArea(c Circle) float64 {
     return math.Pi * c.r*c.r
}
```
lalu di dalam fungsi main kita berikan:
```go
  c := Circle{0, 0, 5}
  fmt.Println(circleArea(c))
```

Satu hal yang perlu diingat adalah bahwa argumen selalu disalin di Go. Jika kami mencoba untuk mengubah salah satu `field` di dalam fungsi `circleArea`, itu tidak akan mengubah variabel asli. Karena itu, kami biasanya menulis fungsi seperti ini:
```go
func circleArea(c *Circle) float64 {
     return math.Pi * c.r*c.r
}
```
Dan ubah pada fungsi main menjadi:
```go
  c := Circle{0, 0, 5}
  fmt.Println(circleArea(&c))
```

## Methods
Meskipun ini lebih baik daripada versi pertama dari kode ini, kita dapat meningkatkannya secara signifikan dengan menggunakan jenis fungsi khusus yang dikenal sebagai `method`:
```go
func (c *Circle) area() float64 {
     return math.Pi * c.r*c.r
}
```
Di antara kata kunci func dan nama fungsi, kami telah menambahkan `receiver`. `receiver` seperti parameter ia memiliki nama dan tipe tetapi dengan membuat fungsi dengan cara ini, kita dapat memanggil fungsi menggunakan `.` operator:
```go
  fmt.Println(c.area())
```
Ini jauh lebih mudah dibaca, kita tidak lagi membutuhkan operator & (Go secara otomatis tahu untuk melewatkan pointer ke lingkaran untuk metode ini) dan karena fungsi ini hanya dapat digunakan dengan `circle`, kita dapat mengganti nama fungsi menjadi `area` saja.

Mari kita lakukan hal yang sama untuk persegi panjang:
```go
type Rectangle struct {
     x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
     l := distance(r.x1, r.y1, r.x1, r.y2)
     w := distance(r.x1, r.y1, r.x2, r.y1)
     return l * w
}
```
dan pada `main` kita memiliki:
```go
r := Rectangle{0, 0, 10, 10}
fmt.Println(r.area())
```

## Embedded Types
`Field struct` biasanya mewakili hubungan. Misalnya `Circle` memiliki `radius`. Misalkan kita memiliki orang struct:
```go
type Person struct {
    Name string
}
func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}
```
Dan kami ingin membuat `struct` `Android` baru. Kita bisa melakukan ini:
```go
type Android struct {
    Person Person
    Model string
}
```

Go mendukung hubungan seperti ini dengan menggunakan tipe yang disematkan. Juga dikenal sebagai bidang `anonim`, jenis yang disematkan terlihat seperti ini:
```go
type Android struct {
    Person
    Model string
}
```
Kita menggunakan tipe (`Person`) dan tidak memberinya nama. Ketika didefinisikan dengan cara ini, struct `Person` dapat diakses menggunakan nama tipe:
```go
  a := new(Android)
  a.Person.Talk()
```

Tetapi kita juga dapat memanggil metode `Person` apa pun secara langsung di `Android`:
```go
a := new(Android)
a.Talk()
```

## Interfaces
Baik dalam kehidupan nyata maupun dalam pemrograman, hubungan seperti ini adalah hal yang lumrah. Go memiliki cara untuk membuat kesamaan yang tidak disengaja ini menjadi eksplisit melalui tipe yang dikenal sebagai `Interface`. Berikut adalah contoh `interface` `Shape`:
```go
type Shape interface {
  area() float64
}
```

Seperti sebuah `struct`, `interface` dibuat menggunakan `type`, diikuti dengan nama dan `interface`. Tapi untuk mendefinisikan `Shape`, kita membutuhkan `"method set"`.Kumpulan `method` adalah daftar `method` yang harus dimiliki suatu tipe untuk "mengimplementasikan" `interface`.

```go
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
```
Kami akan memanggil fungsi ini seperti ini:
```go
  fmt.Println(totalArea(&c, &r))
```

`Interface` juga dapat digunakan sebagai fields:
```go
type MultiShape struct {
     shapes []Shape
}
```
Kita bahkan dapat mengubah `MultiShape` menjadi `Shape` dengan memberinya metode `area`:
```go
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
```
Sekarang `MultiShape` dapat berisi `Circles`, `Rectangles` atau bahkan `MultiShapes` lainnya.

Berikut ini program lebih lengkapnya
```go
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	l, w float64
}

func (r *Rectangle) area() float64 {
	return r.l * r.w
}

type Shape interface {
	area() float64
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{3, 4}
	fmt.Println(totalArea(&c, &r))

	m := MultiShape{}
	m.shapes = append(m.shapes, &c, &r)
	fmt.Println(m.area())
}
```