# Working with GORM (Object Oriented Mapping)

## Dependency yang dibutuhkan
GORM adalah library yang mempermudah koneksi kita ke dalam database sehingga tidak perlu menggunakan query native yang mana kita inisialisasikan query-nya terlebih dahulu tetapi dengan menggunakan ini kita bisa mempermudah dalam operasi dan komunikasi ke dalam database. Beberapa yg dibutuhkan diantaranya.
```go
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
	github.com/go-sql-driver/mysql v1.6.0
```

## Buat Fungsi Koneksi Database
Membuat koneksi disini kita akan lebih di modular-kan yang mana kita akan menyimpan beberapa fungsi berdasarkan kebutuhan tertentu misalnya fungsi koneksi ke database ini kita akan menyimpan code-nya pada folder `internal/database/`.
```go
func NewDatabaseConn() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/recordings?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
	return db
}
```
Karena kita saat ini menggunakan library `gorm` maka kita pun perlu memastikan koneksi database-nya dilakukan dengan benar. Koneksi database menggunakan libary ini terhitung sangat mudah dan applicable sehingga bagi kita tidak perlu ada konfigurasi lebih detail.
```go
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err != nil {
  panic("can't connect to database")
}
```
## Membuat Domain struct
Kita akan mengambil data yang seharusnya kita gunakan dan di simpan sesuai dengan aturan `internal/domain/`.
```go
package domain

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}
```

## Membuat Service and Query ke Database
Selanjutnya kita akan membuat interface dari semua method yang dibutuhkan untuk nantinya akan kita gunakan untuk mengambil, menyimpan, mengubah dan menghapus data.
```go
package album

import (
	"github.com/santekno/using-gorm/internal/domain"
)

type AlbumRepo interface {
	GetAllAlbum() ([]domain.Album, error)
	GetAlbumById(id int64) (domain.Album, error)
	CreateAlbum(album domain.Album) error
	BulkInsertAlbum(albums []domain.Album) error
	UpdateAlbum(album domain.Album) error
	DeleteAlbum(id int) error
}
```

### Inisialisasi Repository module
Untuk membuat suatu fungsi yang berhubungan dengan koneksi ke dalam database maka kita perlu buat suatu modul yang mana nantinya dikumpulkan dalam satu folder `internal/album/repository/` disini kita juga perlu mendeklarasikan koneksi database menggunakan library `gorm` agar bisa digunakan di method-method didalamnya.

```go
type AlbumRepo struct {
	db *gorm.DB
}

func NewAlbumRepo(database *gorm.DB) *AlbumRepo {
	return &AlbumRepo{
		db: database,
	}
}
```

### Buat Fungsi Tambah Bulk
Membuat fungsi tambah atau insert album keseluruhan pada library `gorm` sangat mudah sekali karena kita tidak perlu membuat query string sendiri. Sehingga kita hanya memastikan saja data yg di kirim itu sesuai struct-nya dengan data yg akan diinsert.
```go
func (al *AlbumRepo) BulkInsertAlbum(albums []domain.Album) error {
	err := al.db.CreateInBatches(albums, len(albums))
	if err.Error != nil {
		log.Printf("error when bulk insert err: %v", err)
		return err.Error
	}

	log.Print("Success bulk insert Album")
	return nil
}
```

Penggunaan fungsi `CreateInBatches` ini kita gunakan untuk mengirim data ke database dengan query string yang sudah di buatkan oleh `gorm` sehingga kita tidak perlu lagi mempermasalahkan bagaimana query string-nya ke dalam database. 

## Buat Fungsi Mengambil Semua Data
Pengambilan data menggunakan library `gorm` bisa menggunakan fungsi `Find`. Dalam fungsi ini kita langsung memanggil fungsi tersebut lalu diisi dengan alamat struct `array` domain.Album sehingga data yang diambil lngsung ter-translate ke dalam `struct array` tersebut.

```go
func (al *AlbumRepo) GetAllAlbum() ([]domain.Album, error) {
	var album []domain.Album
	err := al.db.Find(&album).Error
	if err != nil {
		log.Printf("error when getting album err: %v", err)
		return nil, err
	}
	return album, nil
}
```

## Buat Fungsi Mengambil data by Id
Selanjutnya jika kita ingin melakukan filtering pada query biasanya kita menggunakan query string `where`. Nah, didalam library `gorm` juga disediakan fungsi tersebut agar kita bisa menggunakannya lebih ringkas dan mudah.
```go
func (al *AlbumRepo) GetAlbumById(id int64) (domain.Album, error) {
	var album domain.Album
	err := al.db.Where("id = ?", id).Take(&album).Error
	if err != nil {
		log.Printf("error when getting album err: %v", err)
		return domain.Album{}, err
	}
	return album, nil
}
```
Fungsi `Where()` diikuti dengan parameter yang akan di filter akan lebih mudah buat kita jika ingin mengambil data dengan satu baris.

Fungsi `Take()` ini kita gunakan untuk mengambil data yang sudah kita filter lalu di terjemahkan ke dalam struct dalam hal ini kita terjemahkan ke dalam struct Album.

## Buat Fungsi untuk Tambah Album
Pada fungsi tambah Album ini, library `gorm` sudah memberikan fungsi yang mendukung untuk menambahkan data ke dalam database yaitu dengan menggunakan fungsi `Create`.
```go
func (al *AlbumRepo) CreateAlbum(album domain.Album) error {
	tx := al.db.Create(album)
	if tx.Error != nil {
		return tx.Error
	}
	log.Print("Success create Album")
	return nil
}
```

## Buat Fungsi untuk Ubah Album
Pada fungsi ubah Album ini, library `gorm` sudah memberikan fungsi yang mendukung untuk update data ke dalam database yaitu dengan menggunakan fungsi `Save`.

```go
func (al *AlbumRepo) UpdateAlbum(album domain.Album) error {
	var albumUpdate domain.Album
	al.db.Where("id = ?", album.ID).Find(&albumUpdate)
	if albumUpdate.ID == 0 {
		return errors.New("not found")
	}

	albumUpdate.Title = album.Title
	albumUpdate.Artist = album.Artist
	albumUpdate.Price = album.Price

	err := al.db.Save(albumUpdate).Error
	if err != nil {
		return err
	}
	log.Print("Success update Album")
	return nil
}
```
Pada fungsi ini, kita perlu memastikan bahwa data yang dikirim apakah ada dalam database sebelumnya atau tidak ada. Jika ada kita perlu merubah datanya ke dalam struct baru yang nantinya diupdate ke dalam struct tersebut lalu disimpan.

### Buat Fungsi untuk Menghapus Album
Pada fungsi menghapus Album ini, library `gorm` sudah memberikan fungsi yang mendukung untuk menghapus data ke dalam database yaitu dengan menggunakan fungsi `Delete`.
```go
func (al *AlbumRepo) DeleteAlbum(id int) error {
	var albumUpdate domain.Album
	al.db.Where("id = ?", id).Find(&albumUpdate)
	al.db.Delete(albumUpdate)
	log.Print("Success delete Album")
	return nil
}
```

Sebelum kita melakukan `Delete` kita perlu memastikan terlebih dahulu dengan melakukan query dengan fungsi `Where` agar kita tahu bahwa yang akan di delete itu baris yang mana.

## Melakukan inisialisasi module Service pada Main
Akhirnya kita akan melanjutkan bagaimana mengimplementasikan service, repository yang sudah kita definisikan ke dalam `main` program. Ada beberapa bagian yang perlu kita pahami saat kita melakukan inisialisasi pada `main` program.

```go
db := database.NewDatabaseConn()
```
Inisialisai database connection ini dilakukan lebih awal karena variabel `db` ini kita gunakan nantinya ke dalam repository untuk mengakses database dan melakukan query.

```go
app := App{AlbumRepo: repository.NewAlbumRepo(db)}
```
Inisialisai `app` dengan menginputkan repository modul yang dari `interface` struct ke dalam repo inisialisasi yang mana pada deklarasi ini kita mengirim `NewAlbumRepo(db)` sebagai inisialisasi awal repository.

Nah sisanya mari kita tambahkan fungsi-fungsi repository yang sudah kita buat. Misalkan kita akan mengambil semua data album maka kita perlu eksekusi fungsi dibawah ini.
```go
albums, err := app.AlbumRepo.GetAllAlbum()
if err != nil {
  log.Fatal(err)
}
log.Printf("success get all data album : %+v", albums)
```
Lalu jika ingin mengambil data salah satu album saja maka kita tinggal panggil fungsinya seperti ini `app.AlbumRepo.GetAlbumById(1)`.

Jika kamu ingin mengirim data ke dalam database bisa panggil fungsi `app.AlbumRepo.CreateAlbum` atau secara data yang besar bisa pakai `app.AlbumRepo.BulkInsertAlbum()`.

Kita juga sudah bisa menggunakan update data album dan hapus album dengan cara panggil `UpdateAlbum` dan `DeleteAlbum`

Baiklah, selesai sudah semua kita membuat koneksi database menggunakan library `gorm` dalam code memang lebih mudah dan simple karena sudah dihandle oleh library `gorm` untuk membuat query string-nya sehingga kita tidak perlu mendefinisikan.

Tetapi, apakah performance-nya baik?