# Build Service Docker

## Persiapan & Pengenalan Docker
Docker adalah teknologi virtualisasi yang bisa kita gunakan untuk mengisolasi environment service sesuai dengan kebutuhan.

Saat ini kita akan melakukan simulasi dimana service Golang kita akan kita deploy dan build menggunakan Docker. Sebelum lanjut kita perlu beberapa yang harus diinstall yaitu **Docker Application** bisa kita download [disini](https://docs.docker.com/get-docker/).

## Istilah Dalam Docker
### Container
Container adalah sebuah environment ter-isolasi, merupakan bentuk virtualisasi yang lebih kecil dan ringan dibanding VM (Virtual Machine). Virtualisasi pada container disebut dengan Containerization.

### Docker Container
Docker container adalah sebuah container yang di-manage oleh Docker Engine.

### Docker Engine
Docker engine merupakan daemon yang bertugas untuk manajemen container-container.

### Docker Image
Docker Image adalah sebuah file yang di-generate oleh docker, yang file tersebut nantinya digunakan untuk basis pembuatan dan eksekusi container.

### Containerize dan Dockerize
Containerize merupakan istilah terhadap aplikasi yang di-build ke bentuk Image. Sedangkan Dockerize merupakan istilah untuk containerize menggunakan Docker. Perlu diketahui bahwa penyedia container tidak hanya Docker saja, ada banyak engine container lainnya yang bisa dipergunakan.

## Persiapkan Applikasi
Pada sesi ini kita akan mempelajari bagaimana untuk melakukan deploy atau build ke dalam docker environment untuk service kita. Sebelumnya kita harus download terlebih dahulu projek dibawah ini.
```git
git clone https://github.com/santekno/docker-go
```
Setelah kita clone projek-nya lihat pada pada file `Dockerfile` ini merupakan file yang nantinya kita gunakan untuk membuat membuat virtual service menggunakan docker.

Sebelum membahas docker kita pastikan dahulu service kita ini berjalan dengan baik dengan cara jalankan di `local` environment dengan cara dibawah ini.
```bash
go run main.go
```
atau bisa juga kita build terlebih dahulu lalu jalankan binary-nya.
```bash
go build
./docker-go
```
Berikut ini jika project sudah bisa dijalankan dengan baik.
```bash
➜  docker-go git:(main) go run main.go 
GO Docker Tutorial
```

Pada project ini kita hanya menjalankan Rest API satu endpoint yaitu `/hello` yang mana kita menggunakan inisialisasi handler `net/http` bawaan dari Golang Liberary.
```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World")
})

log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
```

## Pembuatan Dockerfile
Pada tahap ini kita akan membuat `image` docker yang nantinya digunakan untuk deploy ke docker local kita. Pastikan docker sudah terinstall ya di komputer kita.

Buat file dengan nama `Dockerfile` ini sebenarnya bisa nama file apa saja tetapi secara umum biasanya memakai nama `Dockerfile`
```bash
touch Dockerfile
```
Lalu kita isi dengan file `Dockerfile` tersebut seperti dibawah ini
```docker
FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/docker-go main.go

ENTRYPOINT ["/app/docker-go"]
```
Berikut adalah penjelasan per baris dari kode di atas.

### FROM golang:alpine
Keyword FROM ini digunakan untuk inisialisasi build stage dan juga menentukan basis Image yang digunakan. Informasi `golang:alpine` di sini adalah basis image yang dimaksud, yaitu image bernama golang dengan tag bernama alpine yang tersedia di laman officil [Docker Hub Golang](https://hub.docker.com/_/golang).

### RUN apk update && apk add --no-cache git
Keyword `RUN` digunakan untuk menjalankan `shell comamnd`. Argument setelahnya, yaitu `apk update && apk add --no-cache git` akan dijalankan di Image `golang:alpine` yang sudah di-set sebelumnya. `Command` tersebut merupakan command `Alpine OS` yang kurang lebih gunanya adalah berikut:
* Command `apk update` digunakan untuk meng-update index packages pada OS.
* Command `apk add --no-cache git` digunakan untuk meng-install Git. 

Kebetulan pada basis image `golang:alpine` by default Git adalah *tidak tersedia*. Jadi harus di-install terlebih dahulu. Git ini nantinya digunakan sewaktu go get dependensi lewat command `go mod tidy`. Meskipun pada contoh project `hello world` tidak menggunakan dependensi eksternal, install saja tidak apa.

### WORKDIR /app
Digunakan untuk menentukan working directory yang pada konteks ini adalah `/app`. Statement ini menjadikan semua statement `RUN` di bawahnya akan dieksekusi pada working directory.

### COPY . .
Digunakan untuk meng-copy file pada argument pertama yaitu `.` yang merepresentasikan direktori yang aktif pada host atau komputer kita (yang isinya file `main.go`, `go.mod`, dan `Dockerfile`), untuk kemudian di-paste ke dalam Image ke working directory yaitu `/app`.

Dengan ini isi `/app` adalah sama persis seperti isi folder project `hello world`.

## RUN go mod tidy
Digunakan untuk validasi dependensi, dan meng-automatisasi proses download jika dependensi yang ditemukan belum ter-download. Command ini akan mengeksekusi go get jika butuh untuk unduh dependensi, makanya kita perlu install Git.

### RUN go build -o /app/docker-go main.go
Command go build digunakan untuk build binary atau executable dari kode program Go. Dengan ini source code dalam working directory akan di-build ke `executable` dengan nama binary.

### ENTRYPOINT ["/app/docker-go"]
Statement ini digunakan untuk menentukan entrypoint container sewaktu dijalankan. Jadi khusus statement `ENTRYPOINT` ini pada contoh di atas adalah yang efeknya baru kelihatan ketika Image di-run ke container. Sewaktu proses build aplikasi ke Image maka efeknya belum terlihat.

Dengan statement tersebut nantinya sewaktu container jalan, maka executable binary yang merupakan aplikasi hello world kita, itu dijalankan di container sebagai entrypoint.

File Dockerfile sudah siap mari kita lanjut ke proses build dan start container.

## Build Image dan Create Container
Dockerfile sudah dibuat saatnya kita melakukan proses build image untuk project kita. 

### Build Images
Pertama masuk ke direktori folder projek, lalu jalankan command docker build berikut.
```bash
cd docker-go
docker build -t docker-go .
```
Command di atas akan melakukan proses `build Image` pada file yang ada di dalam `.` yang merupakan isi folder projek. Projek akan di-build ke sebuah Image dengan nama adalah `docker-go`. Flag `-t` digunakan untuk menentukan nama Image.

Kurang lebih outputnya seperti gambar berikut. Oh iya gunakan command `docker images` untuk menampilkan list semua image yang ada di lokal.
```bash
docker images
```

### Create Container
Selanjutnya kita akan membuat container yang nantinya untuk menjalankan image yang sudah kita buat `build`. Perintahnya bisa dilihat dibawah ini.
```bash
docker container create --name docker-go -e PORT=8080 -p 8080:8080 docker-go
```
Command di atas akan menjalankan sebuah proses yang isinya kurang lebih berikut:
1. Buat container baru dengan nama `docker-go`.
2. Flag `--name` digunakan untuk menentukan nama container.
3. Sewaktu pembuatan container, `env` var `PORT` di-set dengan nilai adalah `8080`.
4. Flag `-e` digunakan untuk menge-set env var. Flag ini bisa dituliskan banyak kali sesuai kebutuhan.
5. Kemudian port `8080` yang ada di luar network docker (yaitu di host/laptop/komputer kita) di map ke port `8080` yang ada di dalam container.
6. Flag `-p` digunakan untuk mapping port antara host dan container. Bagian ini biasa disebut dengan `expose` port.
7. Proses pembuatan container dilakukan dengan Image `docker-go` digunakan sebagai basis image.

Semoga cukup jelas penjabaran di atas. Setelah container berhasil dibuat, cek menggunakan command `docker container ls -a` atau `docker ps -a` untuk menampilkan list semua container baik yang sedang running maupun tidak.

### Start Container
sekarang container juga sudah dibuat, lanjut untuk `start` container tersebut, caranya menggunakan command `docker container start`. Jika sudah, coba cek di browser aplikasi `docker-go`, harusnya sudah bisa diakses.
```bash
docker container start docker-go
docker container ls
```

Bisa dilihat, sekarang aplikasi web `docker-go` sudah bisa diakses dari host/komputer yang aplikasi tersebut running dalam container docker.

Jika mengalami error saat start container, bisa jadi karena port 8080 sudah dialokasikan untuk proses lain. Solusi untuk kasus ini adalah `kill` saja proses yang menggunakan port tersebut, atau `rebuild image` dan `create container` ulang tapi menggunakan port lainnya, selain `8080`.

Pada image di atas juga bisa dilihat penggunaan `command docker container ls` untuk memunculkan list container yang sedand running atau aktif. Untuk menampilkan semua container (aktif maupun non-aktif), cukup dengan menambahkan flag `-a` atau `--all`.

### Stop Container
Untuk stop container bisa dengan command `docker container stop <nama-container-atau-container-id>`.
```bash
docker container stop 
docker container ls
```
### Hapus Container
Untuk hapus container bisa dengan command `docker container rm <nama-container-atau-container-id>`.

docker container rm docker-go
docker container ls

### Hapus Image
Untuk hapus image bisa dengan command `docker image rm <nama-image-atau-image-id>`. Penghapusan image ini harus dipastikan terlebih dahulu tidak ada container yang *running* menggunakan basis image yang ingin dihapus.
```bash
docker image rm my-image-hello-world
docker images
```
## Run Container
Untuk run container sebenarnya ada dua cara ya, yang pertama seperti contoh di atas dengan membuat container nya terlebih dahulu menggunakan command `docker container create` kemudian di start menggunakan command `docker container start`.

Atau bisa juga menggunakan command `docker run`. Command ini akan membuat container baru kemudian otomatis menjalankannya. Tapi kita sampaikan bahwa lewat cara ini tidak ada pengecekan apakah container sudah dibuat atau tidak sebelumnya, pasti akan dibuat container baru.

Mungkin perbandingannya seperti ini:

* Jalankan container lewat create lalu start
```bash
docker container create --name docker-go -e PORT=8080 -p 8080:8080 docker-go
docker container start docker-go
```
* Jalankan container lewat run
```bash
docker container run --name docker-go -e PORT=8080 -p 8080:8080 docker-go
```
Bisa dilihat bukan bedanya, hanya sedikit.

Khusus untuk command `docker run` biasanya dijalankan dengan tambahan beberapa flag agar lebih mudah kontrol-nya, yaitu ditambahkan flag `--rm` dan `-it`.
```bash
docker container run --name docker-go --rm -it -e PORT=8080 -p 8080:8080 docker-go
```

### Flag `--rm`
Flag ini digunakan untuk meng-automatisasi proses penghapusan container sewaktu container tersebut di stop. Jadi kita tidak perlu delete manual pakai `docker container rm`. Hal ini sangat membantu karena command docker run akan membuat container baru setiap dijalankan. Tapi sebenarnya pada contoh sebelumnya kita tidak perlu khawatir akan dibuat container baru karena sudah ada flag `--name`. 
Flag tersebut digunakan untuk menentukan nama container, yang di mana nama container harus unik. Jadi kalau ada duplikasi pasti langsung error. Nah dari sini berarti kalau temen-temen tidak pakai --name sangat dianjurkan paka --rm dalam penerapan docker run.

### Flag `-it`
Flag ini merupakan flag gabungan antara `-i` yang digunakan untuk meng-enable interactive mode dan `-t` untuk enable TTY. Dengan ini kita bisa masuk ke mode interaktif yang di mana jika kita `terminate` atau `kill` command menggunakan `CTRL + C` atau `CMD + C` (untuk mac), maka otomatis container akan di stop.

Nah dengan menggabungkan flag `--rm` dan flag `-it` kita bisa dengan mudah stop kemudian hapus container.

Selain itu ada juga flag yang mungkin penting yaitu `-d` atau *dettach*. Flag ini bisa digabung dengan `-it`. **Dettach** adalah mode di mana ketika command docker run dijalankan, command akan langsung selesai. Dari sini untuk stop container berarti harus menggunakan command docker stop. Contoh:
```bash
docker container run --name docker-go --rm -itd -e PORT=8080 -p 8080:8080 docker-go
docker container stop docker-go
```

## Ringkasan Perintah-perintah yang sering digunakan
```bash
docker build -t <nama-image> .
docker build -t <nama-image>:<version> .
docker build -t <nama-image>:<version> -f <nama-dockerfile> .
docker container create --name <nama-container> -e <environtment-yg-diset> -p <port-dalam>:<port-diluar> <nama-image>
docker container ls -a
docker ps -a
docker images
docker rm <nama-docker>
docker rmi <nama-images>
```