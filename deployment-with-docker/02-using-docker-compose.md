# Using Docker Compose

## Persiapan Docker Compose
Saat ini kita akan melakukan simulasi dimana service Golang kita akan kita deploy dan build menggunakan Docker. Sebelum lanjut kita perlu beberapa yang harus diinstall yaitu **Docker Application** bisa kita download [disini](https://docs.docker.com/get-docker/).

Dan tambahan untuk sesi ini kita akan menggunakan docker-compose, bisa didownload terlebih dahulu [disini](https://docs.docker.com/compose/install/).

## Persiapkan Applikasi
Pada sesi ini kita akan mempelajari bagaimana untuk melakukan deploy atau build ke dalam docker environment untuk service kita. Kita buka folder `postgres-go` yang mana ini sudah kita buat untuk pembuatan service Album.

Sebelum membahas docker kita pastikan dahulu service kita ini berjalan dengan baik dengan cara jalankan di `local` environment dengan cara dibawah ini.
```bash
go run main.go
```
atau bisa juga kita build terlebih dahulu lalu jalankan binary-nya.
```bash
go build
./postgres-go
```
Berikut ini jika project sudah bisa dijalankan dengan baik.
```bash
➜  postgres-go git:(main) ✗ go run main.go         
setting limits
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /hello                    --> main.main.func2 (3 handlers)
[GIN-debug] GET    /albums                   --> github.com/santekno/postgres-go/handler.(*AlbumHandler).GetAllAlbums-fm (3 handlers)
[GIN-debug] GET    /album/:id                --> github.com/santekno/postgres-go/handler.(*AlbumHandler).Get-fm (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

Pada project ini kita hanya menjalankan Rest API satu endpoint yaitu `/hello` yang mana kita menggunakan inisialisasi handler menggunakan library Gin.
```go
router := gin.Default()
router.GET("/hello", func(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H{
    "message": "hello world",
  })
})

router.GET("/albums", application.GetAllAlbums)
router.GET("/album/:id", application.Get)
```

## Pembuatan Dockerfile
Pada tahap ini kita akan membuat `image` docker yang nantinya digunakan untuk deploy ke docker local kita. Pastikan docker sudah terinstall ya di komputer kita.

Buat file dengan nama `Dockerfile` ini sebenarnya bisa nama file apa saja tetapi secara umum biasanya memakai nama `Dockerfile`
```bash
touch Dockerfile
```
Lalu kita isi dengan file `Dockerfile` tersebut seperti dibawah ini
```docker
# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:latest AS build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/album-api main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=build /app/.env /app/.env
COPY --from=build /app/album-api /app/album-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/album-api"]
```
Penjelasan untuk tiap command sudah kita jelaskan di sesi sebelumnya maka kita akan lebih menekankan kepada pembuatan docker-compose.yaml


## Build Images
Pertama masuk ke direktori folder projek, lalu jalankan command docker build berikut.
```bash
cd postgres-go
docker build -t postgres-go .
```
Command di atas akan melakukan proses `build Image` pada file yang ada di dalam `.` yang merupakan isi folder projek. Projek akan di-build ke sebuah Image dengan nama adalah `postgres-go`. Flag `-t` digunakan untuk menentukan nama Image.

Kurang lebih outputnya seperti gambar berikut. Oh iya gunakan command `docker images` untuk menampilkan list semua image yang ada di lokal.
```bash
docker images
```

## Pembuatan `docker-compose.yaml`
Buat file dengan nama `docker-compose.yaml` ini sebenarnya bisa nama file apa saja tetapi secara umum biasanya memakai nama `docker-composer.yaml`
```bash
touch docker-composer.yaml
```
Lalu kita isi dengan file `docker-compose.yaml` tersebut seperti dibawah ini
```yaml
version: '3'

services:
  api:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: album-api
    environment:
      - POSTGRES_URL=postgres://root:secret@db/recordings?sslmode=disable
    ports:
      - "8080:8080"
    depends_on:
      - db
    networks:
      - mynet
  db:
    image: postgres
    container_name: albums-db
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=root
      - POSTGRES_PASSWORD=secret
      - POSTGRES_DB=recordings
    networks:
      - mynet

networks:
  course_network:
    external: true
    name: course_network
  mynet:
    driver: bridge
```
* `services` : Digunakan untuk menginialisasi container apa saja yang akan dibuat, dalam hal ini kita sekarang akan membuat 2 container yaitu container `api` service kita dan `db` untuk penyimpanan database kita.
* `build, context, dockerfile` : Tag ini kita pakai untuk mengarahkan image mana yang akan dijadikan sumber untuk pembuatan container.
* `container_name` : Nama dari container yang akan kita buat
* `environment` : tag ini digunakan untuk mengirim environment config untuk bisa dibaca oleh container API
* `ports` digunakan untuk menginisialisasi port yang akan kita expose ke luar dan yang di expose dari dalam.
* `depends_on` digunakan jika salah satu container memiliki ketergantungan ke container yang lain, sehingga jika container yang lain tidak ada maka tidak bisa dilakukan pembuatan container tersebut.
* `networks` digunakan untuk pemakaian di dalam service tersebut apakah menggunakan network yang sudah diinisialisasi atau disamakan di berbagai container.


## Ringkasan Perintah-perintah yang sering digunakan
```bash
docker-compose up -d
docker-compose up -d --build
docker-compose up -d <nama-service>
docker-compose down
docker-compose down <nama-service>
docker-compose stop <nama-service>
```