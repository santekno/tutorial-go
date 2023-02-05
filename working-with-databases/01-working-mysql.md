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

Selanjutnya pada fungsi `main()` kita akan membagi menjadi beberapa bagian diantaranya yaitu sebagai berikut.

### Inisialisasi Koneksi Database
Pada tahapan ini kita menginisialisasi beberapa konfigurasi yang dibutuhkan untuk membuat koneksi ke dalam database. Diantaranya bisa dilihat sebagai berikut.
```go
cfg := mysql.Config{
    User:   os.Getenv("DBUSER"),
    Passwd: os.Getenv("DBPASS"),
    Net:    "tcp",
    Addr:   "127.0.0.1:3306",
    DBName: "mahasiswa",
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

### Inisialisasi Service Package
Pada tahap ini kita akan membuat folder service untuk memisahkan semua logic query kita ke dalam satu package.
* Buat folder `service`
* Tambahkan file dengan nama file `init.go` dan buatlah isi dari file tersebut seperti dibawah ini.
```go
package services

import "database/sql"

type Services struct {
	db *sql.DB
}

func InitServices(db *sql.DB) Services {
	return Services{
		db: db,
	}
}
```
Maksud pembuatan dari fungsi ini `InitServices` adalah agar kita bisa menggunakan connection yang telah diinisialisasi di `main` proses ke dalam package `services` kita. Sehingga nantinya kita cukup membuat method-method dan menggunakan `db` ini di setiap method.nya.

Jangan lupa ketika sudah melakukan inisialisasi fungsi tersebut maka lakukan juga pemanggilan fungsi tersebut di dalam file `main.go` seperti dibawah ini.
```go
	service := services.InitServices(db)
```

### Mengambil data Semua Mahasiswa dari Database
Pada tahapan selanjutnya kita akan membuat fungsi untuk mengambil data dari database, yang mana pada pertemuan kali ini database yang sudah tersedia adalah semua data mahasiswa.

Fungsi ini akan mengambil data dari database lalu menyimpannya dalam bentuk `struct` yang sudah kita deklarasi sebelumnya.

```go
func (s *Services) GetAllMahasiswa() ([]Mahasiswa, error) {
	var mahasiswas []Mahasiswa

	rows, err := s.db.Query("SELECT id, nama, jenis_kelamin, tempat_lahir, tanggal_lahir, tahun_masuk FROM mahasiswa")
	if err != nil {
		return nil, fmt.Errorf("failed get all mahasiswa %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var mhs Mahasiswa
		if err := rows.Scan(&mhs.ID, &mhs.Nama, &mhs.JenisKelamin, &mhs.TempatLahir, &mhs.TanggalLahir, &mhs.TahunMasuk); err != nil {
			return nil, fmt.Errorf("failed get all mahasiswa %v", err)
		}
		mahasiswas = append(mahasiswas, mhs)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed rows: %v", err)
	}

	return mahasiswas, nil
}
```
Pada fungsi `s.db.Query("SELECT id, nama, jenis_kelamin, tempat_lahir, tanggal_lahir, tahun_masuk FROM mahasiswa")` digunakan untuk mentranslate query dan mengambil data dari database sesuai dengan operasi.

Pada fungsi `rows.Next()` digunakan untuk mem-`fetch` data dari variabel `rows` lalu di translate ke dalam struct `Mahasiswa`. Kita juga perlu melakukan `defer rows.Close()` agar setiap koneksi yang di deklarasi diakhir eksekusi harus di tutup agar terhindar dari `max connection` ke dalam database.

Selanjutnya, jangan lupa juga untuk melakukan pengecekan apakah `err` dan `rows.Err()` memiliki error agar kita tahu bahwa query tersebut terdapat error atau tidak.

### Mengambil data Mahasiswa By ID
Sama halnya dengan mengambil data mahasiswa by id diatas tetapi yang berbeda hanyalah return dari fungsi ini. Berikut ini fungsi untuk mengambil data.
```go
func (s *Services) GetMahasiswaById(id int64) (Mahasiswa, error) {
	var mhs Mahasiswa

	row := s.db.QueryRow("SELECT id,nama,nim,jenis_kelamin,tempat_lahir,tanggal_lahir,tahun_masuk FROM mahasiswa WHERE id = ?", id)
	if err := row.Scan(&mhs.ID, &mhs.Nama, &mhs.NIM, &mhs.JenisKelamin, &mhs.TempatLahir, &mhs.TanggalLahir, &mhs.TahunMasuk); err != nil {
		if err == sql.ErrNoRows {
			return mhs, fmt.Errorf("failed get mahasiswa by id %d: no such mahasiswa", id)
		}

		return mhs, fmt.Errorf("failed get mahasiswa by id %d: %v", id, err)
	}

	return mhs, nil
}
```

Perbedaan dari `get mahasiswa` yaitu pada data mahasiswa disini menggunakan `s.db.QueryRow("SELECT id,nama,nim,jenis_kelamin,tempat_lahir,tanggal_lahir,tahun_masuk FROM mahasiswa WHERE id = ?", id)` yang mana fungsi ini mengembalikan data hanya satu baris.

Pada fungsi ini juga kita menemukan `sql.ErrNoRows` digunakan untuk melakukan pengecekan dan memastikan bahwa data yang diambil itu tidak kosong.

### Menambahkan Mahasiswa
Selanjutnya kita akan menambahkan mahasiswa ke dalam database. Berikut ini fungsi untuk menyimpan data ke dalam database.
```go
func (s *Services) AddMahasiswa(mhs Mahasiswa) (int64, int64, error) {
	result, err := s.db.Exec("INSERT INTO mahasiswa (nama,nim, jenis_kelamin, tempat_lahir, tanggal_lahir, tahun_masuk) VALUES (?, ?, ?, ?, ?, ?)", mhs.Nama, mhs.NIM, mhs.JenisKelamin, mhs.TempatLahir, mhs.TanggalLahir, mhs.TahunMasuk)
	if err != nil {
		return 0, 0, fmt.Errorf("failed add mahasiswa: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, 0, fmt.Errorf("failed add mahasiswa: %v", err)
	}

	sum, err := result.RowsAffected()
	if err != nil {
		return 0, 0, fmt.Errorf("error when getting rows affected")
	}

	return id, sum, nil
}
```

Fungsi `s.db.Exec("INSERT INTO mahasiswa (nama,nim, jenis_kelamin, tempat_lahir, tanggal_lahir, tahun_masuk) VALUES (?, ?, ?, ?, ?, ?)", mhs.Nama, mhs.NIM, mhs.JenisKelamin, mhs.TempatLahir, mhs.TanggalLahir, mhs.TahunMasuk)` berfungsi untuk melakukan query `insert` ke dalam databas sehingga data yang dikirim bisa disimpan ke dalam database.

### Menghapus Mahasiswa
Selanjutnya kita akan menghapus mahasiswa dari dalam database. Berikut ini fungsi untuk menghapus data ke dalam database.
```go
func (s *Services) DeleteMahasiswa(mhsId int64) error {
	if mhsId == 0 {
		return errors.New("mahasiswa ID was zero")
	}

	_, err := s.db.Exec("DELETE FROM mahasiswa WHERE id= ?", mhsId)
	if err != nil {
		log.Printf("error execution : %v", err)
		return err
	}

	return nil
}
```
Perintah untuk menghapus data mahasiswa dalam database sama halnya dengan tambah yaitu menggunakan fungsi `s.db.Exec` hanya yang membedakan yaitu pada query yang digunakan saja yaitu `DELETE FROM mahasiswa WHERE id=?`.

### Menambahkan Mahasiswa menggunakan Transaction Batching
Biasanya kita kadang membutuhkan operasi untuk menyimpan data mahasiswa secara bulk (langsung banyak) agar mengefisienkan waktu saat pengisian data dibandingkan dengan satu persatu mengisi data mahasiswanya. Maka kita perlu membuatkan suatu method yang bisa mendukung batching data ke mahasiswa ke dalam database. Berikut ini bagaimana kita membuat method khusus untuk batching.
```go
func (s *Services) BulkInsertUsingTransaction(mahasiswas []Mahasiswa) ([]int64, error) {
	var insertID []int64

	if len(mahasiswas) == 0 {
		return insertID, errors.New("mahasiswa record was empty")
	}

	tx, err := s.db.Begin()
	if err != nil {
		return insertID, errors.New("begin mahasiswa transaction error")
	}

	defer tx.Rollback()

	for _, mhs := range mahasiswas {
		result, err := tx.Exec("INSERT INTO mahasiswa (nama, nim, jenis_kelamin, tempat_lahir, tanggal_lahir, tahun_masuk) VALUES (?, ?, ?, ?, ?, ?)", mhs.Nama, mhs.NIM, mhs.JenisKelamin, mhs.TempatLahir, mhs.TanggalLahir, mhs.TahunMasuk)
		if err != nil {
			log.Printf("error execution : %v", err)
			continue
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			log.Printf("error last insert : %v", err)
		}

		insertID = append(insertID, lastInsertId)
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error commit : %v", err)
		return insertID, err
	}

	return insertID, err
}
```

Ada beberapa catatan saat kita menggunakan `transaction` database diantaranya
* Pada awal method menggunakan `s.db.Begin()` ini dimaksudkan agar kita menginisialisasi proses transaksi ke dalam database yang mana pada saat ini kita mengalokasikan koneksi database khusus untuk transaksi ini.
* Penggunaan `defer tx.Rollback()` digunakan agar saat ada data ditengah atau dibagian data tertentu ada yang membuat error sehingga data tidak masuk ke dalam database, maka tiap transaksi tersebut akan dikembalikan ke semula atau biasanya disebut `rollback`.
* Penggunaan `tx.Commit()` digunakan untuk mengakhiri semua proses transaksi ke dalam databas sehingga semua data akan langsung disimpan ke dalam database.

Sudah pahamkah kita bagaimana mengoperasikan semua dan mengkomunikasikan datanya ke dalam database?
Semoga saja teman-teman bisa memahami semua yang sudah dijelaskan pertahapannya di tutorial ini.