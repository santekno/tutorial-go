# Working Postgres Database

## Dependency
Pastikan sudah buat projek dengan `go mod init postgres-go` pada folder `postgres-go` dan dependency yang akan kita pakai yaitu menggunakan
```go
import _ "github.com/lib/pq"
```
Jika belum ada bisa kita coba download terlebih dahulu dependency diatas dengan cara perintah dibawah ini.
```bash
go get github.com/lib/pg
```

## Buat Fungsi Koneksi Database
Pada pembuatan fungsi koneksi database ini kita akan menggunakan fungsi dari sql yang mana sudah disiapkan oleh library `go`. Mari kita buat fungsi koneksi database ini dengan mengembalikan return `db` connection.
```go
func openDB(dsn string, setLimits bool) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if setLimits {
		fmt.Println("setting limits")
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
```
Pada fungsi koneksi ini kita menggunakan `sql.Open` untuk melakukan koneksi database `Postgres` yang mana kita juga menerima parameter dari fungsi utama itu `dsn` yang berisi server lokasi, user, password dan database yang dituju.
```go
db, err := sql.Open("postgres", dsn)
```

Lalu kita juga memberikan beberapa konfigurasi untuk memastikan koneksi ini memiliki `SetMaxOpenConns` Koneksi saat kita jalankan service ini dan juga `SetMaxIdleConns` kita set untuk memastikan koneksi kita tidak salah konfigurasi dan berjalan dengan normal.
```go
db.SetMaxOpenConns(5)
db.SetMaxIdleConns(5)
```

Selanjutnya kita pastikan juga buat context set limit timeout yang mana disini berguna untuk memastikan koneksi ke database ini tidak terlalu lama. Dalam hal ini kita set 5 second batas maksimal timeout-nya.
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

Dan terakhir yaitu kita melakukan `PingContext` ini memastikan bahwa koneksi ke dalam database ini tidak terhambat oleh apapun misalkan delay atau network.
```go
err = db.PingContext(ctx)
```

## Membuat Service and Query ke Database
Pada tahapan ini, kita akan membuat beberapa fungsi untuk mengakses data dari database fungsinya pun kita buat dengan enkapsulasi `interface`. Mari kita langsung jalankan dan praktikan. Buat folder `services` lalu kita buat juga file didalam folder tersebut dengan nama file `contract.go`.
```go
package services

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

type AlbumService interface {
	Get(id int64) (*Album, error)
	Create(album *Album) error
	GetAllAlbum() ([]Album, error)
	BatchCreate(albums []Album) error
	Update(Album Album) error
	Delete(id int64) error
}
```
Bisa kita lihat bahwa file `contract` tersebut berisi tentang data `struct` yang mana akan digunakan untuk dasar struktur datanya `album`. 

Sedangkan fungsi-fungsi yang akan kita gunakan yaitu kebutuhan dari tambah, ubah, hapus. Berikut ini kita buat `interface` yanng nantinya akan kita implementasi juga pada fungsi-fungsi tersebut.
```go
type AlbumService interface {
	Get(id int64) (*Album, error)
	Create(album *Album) error
	GetAllAlbum() ([]Album, error)
	BatchCreate(albums []Album) error
	Update(Album Album) error
	Delete(id int64) error
}
```

### Mengamnbil Data Album by Id
Mari kita buat fungsi yang nantinya bisa mengambil data dari database postgres. Sebelum ke fungsi yang terkait kita perlu definisikan terlebih dahulu fungsi-fungsi yang ada dengan menyertakan koneksi database yang sudah kita definisikan di awal.
```go
type PostgresService struct {
	db *sql.DB
}

func NewPostgresService(db *sql.DB) *PostgresService {
	return &PostgresService{db: db}
}
```

Inisialisasi ini kita gunakan agar bisa kita membuat fungsi yang mengambil data ke dalam database itu dengan koneksi yang sudah ada. Lalu selanjutnya mari kita buat fungsi dari mengambil data album mengunakan parameter `id`.
```go
func (p *PostgresService) Get(id int64) (*Album, error) {
	query := `
        SELECT id, title, artist, price
        FROM public.album
        WHERE id = $1`

	var album Album

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&album.ID,
		&album.Title,
		&album.Artist,
		&album.Price,
	)

	if err != nil {
		return nil, err
	}

	return &album, nil
}
```

Pada fungsi diatas itu kita melihat buat query ke dalam database seperti dibawah ini.
```go
query := `
        SELECT id, title, artist, price
        FROM public.album
        WHERE id = $1`
```

Lalu kita juga pastikan bahwa query tersebut perlu `timeout` agar tidak terlalu lama mengambil data ke dalam database. Pada program ini kita menggunakan timeout 15 second.
```go
ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
defer cancel()
```

Selanjutnya kita juga menggunakan perintah `QueryRowContext` untuk mengambil datanya setelah kita inisialisasi koneksi ke database.
```go
err := p.db.QueryRowContext(ctx, query, id).Scan(
		&album.ID,
		&album.Title,
		&album.Artist,
		&album.Price,
	)
```

Perintah `QueryRowContext` ini digunaakn untuk mengambil data dari database dengan mengembalikan hanya satu baris data saja. Selanjutnya jika datanya ada maka akan kita simpan data tersebut ke dalam variabel `album`.

### Menyimpan Data Album ke Database
Selain kita mengambil data dari database, kita juga harus bisa mengirim data dari sistem ke dalam database. Berikut ini kita akan membuat fungsi `create`.
```go
func (p *PostgresService) Create(album *Album) error {
	query := `
        INSERT INTO public.album (id, title, artist, price) 
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	return p.db.QueryRowContext(ctx, query, album.Title, album.Artist, album.Price).Scan(&album.ID)
}
```

Seperti biasa karena ini termasuk ke dalam native query yang mana kita menentukan query sendiri didalam kode maka kita perlu definisikan terlebih dahulu query-nya. Berikut ini query untuk insert data ke dalam database.
```go
query := `
  INSERT INTO public.album (id, title, artist, price) 
  VALUES ($1, $2, $3, $4)
  RETURNING id`
```

Lakukan set timeout agar lebih terjaga dan selanjutnya kita execute create disini kita menggunakan query sama seperti mengambil data karena kita akan mengambil id setelah ditambahkan.
```go
return p.db.QueryRowContext(ctx, query, album.Title, album.Artist, album.Price).Scan(&album.ID)
```

### Mengambil Semua Data dalam Database
Mari kita buat fungsi selanjutnya yaitu mengambil semua data dari database sebagai berikut.
```go
func (p *PostgresService) GetAllAlbum() ([]Album, error) {
	query := `
		SELECT id, title, artist, price
		FROM album`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var albums []Album

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album Album
		err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.Artist,
			&album.Price,
		)
		if err != nil {
			return nil, err
		}

		albums = append(albums, album)
	}

	return albums, nil
}
```

Pada fungsi ini kita akan mengambil semua datanya dan seperti kita lihat fungsi query yang digunakan berbeda dengan query sebelumnya, yaitu kita menggunakan query `QueryContext` yang mana hasil dari kembalian fungsi ini memiliki multi rows sehingga kita perlu melakukan perulangan untuk menyimpan datanya ke dalam array struct.
```go
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
```

Berikut ini `QueryContext` digunakan untuk mengambil data dan jangan lupa karena kembalian ini memiliki multiple rows maka kita perlu juga dilakukan close connection tiap rows-nya diakhir fungsi. Kalau disini kita pakai `defer` agar dijalankan setelah kembali ke fungsi utama.
```go
defer rows.Close()
```

Selanjutnya kita akan mengambil data dari `rows` untuk dipindahkan ke dalam satu struct array.
```go
for rows.Next() {
  var album Album
  err := rows.Scan(
    &album.ID,
    &album.Title,
    &album.Artist,
    &album.Price,
  )
  if err != nil {
    return nil, err
  }

  albums = append(albums, album)
}
```

Dan diakhiri dengan perintah `return albums, nil` pada fungsinya.

### Membuat Batch Tambah Album
Pada fungsi ini kita akan membuat batching data yang mana kia mengirim data dengan jumlah lebih dari satu. Ini bisa kita gunakan untuk data-data yang kita kirim sekaligus. Mari kita lihat fungsinya yang akan kita buat.
```go
func (p *PostgresService) BatchCreate(albums []Album) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)`

	for _, album := range albums {
		_, err := tx.ExecContext(ctx, query, album.Title, album.Artist, album.Price)
		if err != nil {
			log.Printf("error execute insert err: %v", err)
			continue
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
```

Fungsi batch ini kita menggunakan beberapa perintah batching transaksi ke dalam database. Pada tahapan awal kita akan membuka transaksi ke db dengan perintah dibawah ini.
```go
tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
```

Fungsi `Begin()` digunakan untuk membuka transaksi dan diakhiri dengan `defer tx.Rollback()` jika di tengah-tengah terjadi sesuatu yang mengakibatkan `error`.

Selanjutnya kita lakukan perulangan untuk melakukan insert data ke dalam database.
```go
for _, album := range albums {
  _, err := tx.ExecContext(ctx, query, album.Title, album.Artist, album.Price)
  if err != nil {
    log.Printf("error execute insert err: %v", err)
    continue
  }
}
```

Saat melakukan `create` data ke dalam database kita menggunakan fungsi `ExecContext` yang mana berguna untuk menyimpan sementara di dalam database. Dan diakhiri dengan `Commit()` untuk menyimpan secara permanen ke dalam database.
```go
err = tx.Commit()
if err != nil {
  return err
}
```

### Melakukan Ubah Data Album
Melakukan perubahan dari satu data kita perlu membuat fungsi yang mengirimkan parameter `id` untuk menyeleksi data yang akan diubah. Berikut ini fungsi lengkapnya untuk mengubah album.
```go
func (p *PostgresService) Update(album Album) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `UPDATE album set title=$1, artist=$2, price=$3 WHERE id=$4`
	result, err := p.db.ExecContext(ctx, query, album.Title, album.Artist, album.Price, album.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected update : %d", rows)
	return nil
}
```
Perintah `ExecContext` digunakan sama halnya dengan tambah karena operasi ini bisa kita gunakan asalkan query-nya untuk kebutuhan update dalam hal ini kita gunakan untuk update data.
```go
query := `UPDATE album set title=$1, artist=$2, price=$3 WHERE id=$4`
result, err := p.db.ExecContext(ctx, query, album.Title, album.Artist, album.Price, album.ID)
if err != nil {
  return err
}
```

Perintah selanjutnya yaitu memastikan apakah data tersebut berhasil berubah atau tidak. Disini kita menggunakan fungsi `RowsAffected()` yang memastikan apakah terdapat baris data yang sudah terupdate setelah operasi ini atau tidak.
```go
rows, err := result.RowsAffected()
if err != nil {
  return err
}
```

## Menghapus data Album
Operasi atau fungsi untuk menghapus data pada program ini kita menghapus satu persatu berdasarkan input parameter dari fungsi tersebut. Berikut fungsi lengkapnya.
```go
func (p *PostgresService) Delete(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	query := `DELETE from album WHERE id=$1`
	result, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("Affected delete : %d", rows)
	return nil
}
```

Pada fungsi ini kita memanggil `ExecContext` untuk menghapus data ke dalam database. Hal ini digunakan sama halnya seperti update dan insert.
```go
query := `DELETE from album WHERE id=$1`
result, err := p.db.ExecContext(ctx, query, id)
if err != nil {
  return err
}
```

Dilanjutkan juga dengan perintah `RowsAffected()` agar memastikan data yang akan kita delete itu sukses atau tidak ditemukan pada database.
```go
rows, err := result.RowsAffected()
if err != nil {
  return err
}
```

## Melakukan inisialisasi module Service pada Main
Setelah kita membuat program service yang mana kita membuat fungsi masing-masih untuk operasi ke dalam database saatnya kita mengintegrasikan semua tersebut ke dalam `main` program. Berikut ini main fungsi selengkapnya.
```go
type app struct {
	AlbumService services.AlbumService
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := openDB(os.Getenv("POSTGRES_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)

	application := app{AlbumService: services.NewPostgresService(db)}

	albums, err := application.AlbumService.GetAllAlbum()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("all album : %v\n", albums)

	albumNo1, err := application.AlbumService.Get(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("album number 1 : %v\n", albumNo1)

	err = application.AlbumService.BatchCreate([]services.Album{
		{Title: "Hari Yang Cerah", Artist: "Peterpan", Price: 50000},
		{Title: "Sebuah Nama Sebuah Cerita", Artist: "Peterpan", Price: 50000},
		{Title: "Bintang Di surga", Artist: "Peterpan", Price: 60000},
	})
	if err != nil {
		log.Fatal(err)
	}

	albumNo1.Price = 70000
	err = application.AlbumService.Update(*albumNo1)
	if err != nil {
		log.Fatal(err)
	}

	err = application.AlbumService.Delete(albumNo1.ID)
	if err != nil {
		log.Fatal(err)
	}
}
```

Mari kita bahas satu per satu perintah apa saja yang di jalankan pada fungsi main. 

### Mengambil konfigurasi dari file .env
Pada program main ini kita menyimpan semua konfigurasi koneksi ke dalam database disimpan pada file terpisah misalnya kita menyimpan di dalam file `.env`.
```go
err := godotenv.Load(".env")
if err != nil {
  log.Fatalf("Error loading .env file")
}
```

Perintah diatas digunakan untuk meload data konfigurasi dalam file `.env`. lalu disimpan ke dalam env server atau komputer yang sedang dijalankan. Biasanya untuk mengambil data dari environment tersebut bisa dengan perintah seperti ini.
```go
os.Getenv("<nama-konfigurasi>")
```

Dalam hal ini kita menyimpan koneksi URL ke dalam database yang disimpan ke dalam environment sehingga kita perlu mengambil data konfigurasinya itu menggunakan perintah `os.Getenv("POSTGRES_URL")`.

Selanjutnya kita inisialisasi koneksi ke dalam database.
```go
db, err := openDB(os.Getenv("POSTGRES_URL"), true)
if err != nil {
  log.Fatalln(err)
}
```

Nah saat ini kita perlu juga memastikan database koneksi ini perlu diakhiri `close`. Bisa juga kita gunakan menggunakan `defer`.
```go
defer func(db *sql.DB) {
  err := db.Close()
  if err != nil {
    log.Fatalln(err)
  }
}(db)
```

Agar `service` kita melakukan akses pada main program maka kita perlu melakukan inisialisasi module `service`-nya. 
```go
application := app{AlbumService: services.NewPostgresService(db)}
```

Sisanya adalah pemanggilan fungsi-fungsi `service` yang sudah kita definisikan pada fungsi.