# Working with MySQL

## Package atau Library
```go
import  "github.com/go-sql-driver/mysql"
```

## Inisialisasi Projek
Siapkan folder baru dengan nama `mysql-native`, lalu buat inisialisasi module golang agar lebih termodular. Berikut perintah cepatnya.
```bash
$ mkdir mysql-native
$ cd mysql-native
$ go mod init github.com/santekno/mysql-native
```

## Tambah dependency
Setelah module dibuat kita perlu juga untuk menambahkan dependency `mysql` yang mana pada tutorial ini kita akan menggunakan dependency `"github.com/go-sql-driver/mysql"`. Tambahkan dependency tersebut pada projek modul ini dengan cara perintah dibawah ini.
```bash
➜  mysql-native git:(main) ✗ go get github.com/go-sql-driver/mysql 
go: downloading github.com/go-sql-driver/mysql v1.6.0
go get: added github.com/go-sql-driver/mysql v1.6.0
```
Lihat file di `go.mod` jika sudah terdapat dependency seperti ini berarti dependency tersebut sudah kita pasang.
```bash
➜  mysql-native git:(main) ✗ cat go.mod   
module github.com/santekno/mysql-native

go 1.17

require github.com/go-sql-driver/mysql v1.6.0 // indirect
```
Selanjutnya lakukan perintah dibawah ini untuk mendownload dependency-nya pada folder `vendor/`.
```bash
➜  mysql-native git:(main) ✗ go mod vendor
```

## Pembuatan program
Pada saat ini kita akan membuat program dalam 1 file `main` saja belum melakukan beberapa teknik struktural yang mengorganisasikan beberapa file folder atau sering kita sebut dengan `Framework`. 

Pada program ini sederhana sehingga kita hanya perlu membutuhkan 1 file `main.go` saja untuk mengoperasikan semua yang akan kita lakukan.

Buat file `main.go` pada projek dan program tersebut kita perlu inisialisasi database connection yang bisa dilihat seperti dibawah ini.
```go
var db *sql.DB
```
Lalu berikutnya kita akan menyimpan data tersebut dalam `struct` seperti ini.
```go
type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}
```
Selanjutnya pada fungsi `main()` kita akan membagi menjadi beberapa bagian diantaranya yaitu sebagai berikut.

### Inisialisasi Koneksi Database
Pada tahapan ini kita menginisialisasi beberapa konfigurasi yang dibutuhkan untuk membuat koneksi ke dalam database. Diantaranya bisa dilihat sebagai berikut.
```go
cfg := mysql.Config{
    User:   os.Getenv("DBUSER"),
    Passwd: os.Getenv("DBPASS"),
    Net:    "tcp",
    Addr:   "127.0.0.1:3306",
    DBName: "recordings",
}
```
Pada dependency `mysql` ini ada beberapa konfigurasi yang lebih lengkap bisa kita lihat pada dokumentasi dari dependency tersebut.
```go
type Config struct {
	User             string            // Username
	Passwd           string            // Password (requires User)
	Net              string            // Network type
	Addr             string            // Network address (requires Net)
	DBName           string            // Database name
	Params           map[string]string // Connection parameters
	Collation        string            // Connection collation
	Loc              *time.Location    // Location for time.Time values
	MaxAllowedPacket int               // Max packet size allowed
	ServerPubKey     string            // Server public key name
	pubKey           *rsa.PublicKey    // Server public key
	TLSConfig        string            // TLS configuration name
	tls              *tls.Config       // TLS configuration
	Timeout          time.Duration     // Dial timeout
	ReadTimeout      time.Duration     // I/O read timeout
	WriteTimeout     time.Duration     // I/O write timeout

	AllowAllFiles           bool // Allow all files to be used with LOAD DATA LOCAL INFILE
	AllowCleartextPasswords bool // Allows the cleartext client side plugin
	AllowNativePasswords    bool // Allows the native password authentication method
	AllowOldPasswords       bool // Allows the old insecure password method
	CheckConnLiveness       bool // Check connections for liveness before using them
	ClientFoundRows         bool // Return number of matching rows instead of rows changed
	ColumnsWithAlias        bool // Prepend table alias to column names
	InterpolateParams       bool // Interpolate placeholders into query string
	MultiStatements         bool // Allow multiple statements in one query
	ParseTime               bool // Parse time values to time.Time
	RejectReadOnly          bool // Reject read-only connections
}
```
Pada program ini kita hanya melakukan konfigurasi dasar saja diantaranya yaitu
| Konfigurasi | Informasi |
| :- | :--------- |
| `User` | user koneksi yang dibutuhkan untuk ke database mysql |
| `Passwd` | passowrd untuk user yang dibutuhkan untuk koneksi ke database mysql |
| `Net` | protokol yang digunakan untuk koneksi ke database |
| `Addr` | alamat yang menunjukkan server dari database |
| `DBName` | nama dari database yang dituju |

### Membuat file .env
Kita lihat bahwa untuk koneksi database ini membutuhkan `user` dan `password` yang akan diambil dari environtment. Dilihat dari pemanggilan fungsi `os.Getenv("DBUSER")` dan `os.Getenv("DBPASS")` yang mana ini berfungsi untuk mengambil variable dari environment untuk mendapatkan `user` dan `password` dari database yang akan terkoneksi.

Pada saat ini kita bahas, untuk mendapatkan environment ini biasanya digunakan untuk memisahkan beberapa variabel global yang mana ini dibutuhkan agar lebih `configurable` jika kita sudah live dalam menjalankan program-nya. 

Bagaimana caranya? Kita perlu buat file baru `.env` lalu isi file tersebut dengan kode seperti ini.
```env
DBUSER=<username-database>
DBPASS=<password-database>
```

### Pengecekan Koneksi dan Ping
Kita akan menanjutkan kembali untuk melengkapi program pada fungsi `main()`. Setelah kita mengisi konfigurasi yang dibutuhkan untuk koneksi ke database `mysql`. Saatnya kita perlu memanggil koneksi dan pengetesan apakah koneksi tersebut bisa berjalan atau tidak.

Berikut ini program untuk melakukan koneksi database dan pengetesan koneksinya.
```go
var err error
db, err = sql.Open("mysql", cfg.FormatDSN())
if err != nil {
  log.Fatal(err)
}

pingErr := db.Ping()
if pingErr != nil {
  log.Fatal(pingErr)
}
fmt.Println("Connected!")
```
Pada fungsi `sql.Open("mysql",cfg.FormatDNS())` ini digunakan untuk melakukan koneksi ke dalam database, jika koneksi ini tidak bisa dilakukan maka fungsi ini pun mengeluarkan `err` yang kita tangkap sehingga program akan error karena tidak bisa melakukan koneksi ke `database`.

Lalu fungsi `db.Ping()` ini digunakan untuk memastikan bahwa koneksi tersebut sudah bisa digunakan untuk mengambil, menyimpan, bahkan melakukan penghapusan data ke dalam database.

### Mengambil data Album By Artis dari Database
Pada tahapan selanjutnya kita akan membuat fungsi untuk mengambil data dari database, yang mana pada pertemuan kali ini database yang sudah tersedia adalah data album berdasarkan artis.

Fungsi ini akan mengambil data dari database lalu menyimpannya dalam bentuk `struct` yang sudah kita deklarasi sebelumnya.
```go
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	
  if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}
```
Pada fungsi `db.Query("SELECT * FROM album WHERE artist = ?", name)` digunakan untuk mentranslate query dan mengambil data dari database sesuai dengan operasi.

Pada fungsi `rows.Next()` digunakan untuk mem-`fetch` data dari variabel `rows` lalu di translate ke dalam struct `Album`.

Kita juga perlu melakukan `defer rows.Close()` agar setiap koneksi yang di deklarasi diakhir eksekusi harus di tutup agar terhindar dari max connection ke dalam database.

Selanjutnya, jangan lupa juga untuk melakukan pengecekan apakah `err` dan `rows.Err()` memiliki error agar kita tahu bahwa query tersebut terdapat error atau tidak.

### Mengambil data Album
Sama halnya dengan mengambil data album by artis diatas tetapi yang berbeda hanyalah return dari fungsi ini. Berikut ini fungsi untuk mengambil data album.
```go
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}
```
Perbedaan dari `albumsByArtist` yaitu pada data album disini menggunakan `db.QueryRow("SELECT * FROM album WHERE id = ?", id)` yang mana fungsi ini mengembalikan data hanya satu baris.

Pada fungsi ini juga kita menemukan `sql.ErrNoRows` digunakan untuk melakukan pengecekan dan memastikan bahwa data yang diambil itu tidak kosong.

### Menambahkan Album
Selanjutnya kita akan menambahkan album ke dalam database. Berikut ini fungsi untuk menyimpan data ke dalam database.
```go
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
```
Fungsi `db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)` berfungsi untuk melakukan query `insert` ke dalam databas sehingga data yang dikirim bisa disimpan ke dalam database.