# Working MongoDB Database

# Dependecy yang dibutuhkan
Menambahkan beberapa dependency yang digunakan
```go
  "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
```

## Membuat Koneksi Database
membuat fungsi koneksi database ke dalam mongodb.
```go
func connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb+srv://test:test@cluster0.awkve.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return client.Database("recordings"), nil
}
```

Berbeda dengan yang lain bahwa untuk melakukan koneksi ini kita perlu inisialisasi koneksi selanjutnya kita buat koneksi ke dalam database `mongodb`.
```go
clientOptions := options.Client()
clientOptions.ApplyURI("mongodb+srv://test:test@cluster0.awkve.mongodb.net/?retryWrites=true&w=majority")
```

Lalu dilanjutkan dengan membuat client inisialisasi digunakan agar client memastikan koneksi ke dalam database mongodb.
```go
client, err := mongo.NewClient(clientOptions)
if err != nil {
  return nil, err
}
```

Selanjutnya kita buat `connect` agar kita pastikan koneksinya ke dalam database.
```go
err = client.Connect(context.Background())
if err != nil {
  return nil, err
}
```

Diakhiri dengan memilih database (collection) mana yang akan kita pilih.
```go
return client.Database("recordings"), nil
```

## Membuat fungsi Tambah data ke dalam Album
Pada fungsi tambah data kita perlu memanggil koneksi terlebih dahulu agar kita bisa mengambil data colection dan insert data ke dalam database.

```go
func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()

	_, err = db.Collection("album").InsertOne(ctx, Album{ID: 1, Title: "Hari Yang Cerah", Artist: "Peterpan", Price: 50000})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("album").InsertOne(ctx, Album{ID: 2, Title: "Sebuah Nama Sebuah Cerita", Artist: "Peterpan", Price: 50000})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}
```
Perintah untuk insert ini digunakan yaitu `InsertOne` yang mana kita menyimpan satu data collection ke dalam database `recordings`.
```go
_, err = db.Collection("album").InsertOne(ctx, Album{ID: 1, Title: "Hari Yang Cerah", Artist: "Peterpan", Price: 50000})
if err != nil {
  log.Fatal(err.Error())
}
```

## Membuat Menampilkan album 
Selanjutnya kita akan membuat fungsi mengambil data ke dalam mongodb untuk dengan `id` = `. Berikut fungsi lengkapnya.
```go
func find() {
	ctx := context.Background()
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("album").Find(ctx, bson.M{"id": 1})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]Album, 0)
	for csr.Next(ctx) {
		var row Album
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Title  :", result[0].Title)
		fmt.Println("Artist :", result[0].Artist)
		fmt.Println("Price  :", result[0].Price)
	}
}
```
Saat mengambil data ke dalam mongodb kita mengambil collection `album` lalu kita pakai fungsi `Find` untuk mengambil salah satu data yang amna disini kita gunakan adalah id nomor satu.
```go
csr, err := db.Collection("album").Find(ctx, bson.M{"id": 1})
if err != nil {
  log.Fatal(err.Error())
}
defer csr.Close(ctx)
```

Jangan lupa setiap conection db kita tutup dengan memanggil perintah dibawah ini.
```go
defer csr.Close(ctx)
```

Selanjutnya kita akan menyimpan data hasil query ke dalam variable yang nantinya kita kirim ke dalam main fungsi.
```go
result := make([]Album, 0)
for csr.Next(ctx) {
  var row Album
  err := csr.Decode(&row)
  if err != nil {
    log.Fatal(err.Error())
  }

  result = append(result, row)
}

if len(result) > 0 {
  fmt.Println("Title  :", result[0].Title)
  fmt.Println("Artist :", result[0].Artist)
  fmt.Println("Price  :", result[0].Price)
}
```
## Menampilkan data keseluruhan Album
Pada fungsi ini kita ingin mengambil semua data yang ada pada album collection agar bisa ditampilkan secara keseluruhan.

```go
func findall() {
	ctx := context.Background()
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("album").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]Album, 0)
	for csr.Next(ctx) {
		var row Album
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		for _, res := range result {
			fmt.Println("Title  :", res.Title)
			fmt.Println("Artist :", res.Artist)
			fmt.Println("Price  :", res.Price)
		}
	}
}
```

Kita akan memanggil fungsi yang sama seperti yang diatas tetapi output fungsi pada fungsi ini mengeluarkan array struct `Album`
## Melakukan Ubah data Album
Pada saat kita ingin mengubah data album disini perlu memilih id mana yang akan kita update agar bisa diterima. Berikut ini fungsi untuk melakukan update pada collection database mongodb.
```go
func update() {
	ctx := context.Background()
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"id": 2}
	var changes = Album{ID: 2, Title: "Bintang Di surga", Artist: "Peterpan", Price: 60000}
	_, err = db.Collection("album").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Update success!")
}
```
Berbeda dengan query yang diatas, untuk fungsi update kita perlu melakukan set selector agar kita tahu mana yang akan kita update datanya. Lalu dilanjutkan dengan changes yang akan kita update

```go
var selector = bson.M{"id": 2}
var changes = Album{ID: 2, Title: "Bintang Di surga", Artist: "Peterpan", Price: 60000}
```
Setelah itu kita gunakan fungsi `UpdateOne` untuk melakukan update ke dalam database mongodb.
```go
_, err = db.Collection("album").UpdateOne(ctx, selector, bson.M{"$set": changes})
if err != nil {
  log.Fatal(err.Error())
}
```


## Melakukan Hapus Data Album
Fungsi selanjutnya yaitu kita menambahkan fungsi hapus album. Dalam hal ini kita ingin menghapus data koleksi id 2 dengan fungsi lengkapnya dibawah ini.
```go
func remove() {
	ctx := context.Background()
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"id": 2}
	_, err = db.Collection("album").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Remove success!")
}
```
Seperti halnya update, kita juga perlu menyelesi mana yang akan kita update contoh disini yaitu menggunakan `bson.M{"id", 2}` yang akan kita hapus. Selanjutnya kita eksekusi dengan fungsi `DeleteOne`.
```go
var selector = bson.M{"id": 2}
_, err = db.Collection("album").DeleteOne(ctx, selector)
if err != nil {
  log.Fatal(err.Error())
}
```

Baiklah, semua fungsi sudah kita definisikan tingal kita memanggil tiap fungsi yang sudah kita definisikan ke dalam main fungsi yang mana akan kita eksekusi ke depannya.