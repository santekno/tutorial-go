# Membuat Versi Mayor dalam Projek

Dalam suatu projek ada saatnya kita harus melakukan update beberapa fungsi atau terkait flow bisnis dalam program tersebut agar sesuai dengan kebutuhan yang akan datang. Maka perlu ada penataan program yang terstruktur agar semua program yang terikat dengan dependency tidak menghasilkan program menjadi `error` dan tidak terkendali. 

Maka fungsi dari `version` adalah memperbarui tanpa harus semua dependency yg terikat harus melakukan juga migrasi ke program yang terbaru. *So*, mari kita coba lakukan `update` versi mayor dari project `hello`.

Pertama, kita buat folder `v2` sesuai dengan aturan go bahwa `versioning` harus berbeda folder. Setelah itu kita salin semua program `.go` ke dalam folder `v2`
```bash
$ mkdir v2
$ cp *.go v2/
```

Sekarang, mari kita buat berkas `go.mod` untuk v2 dengan menyalin berkas `go.mod` yang sudah ada dan menambahkan *suffix* `v2` ke path modul:
```bash
$ cp go.mod v2/go.mod
$ go mod edit -module example.com/hello/v2 v2/go.mod
```

Ingatlah bahwa versi v2 diperlakukan sebagai modul terpisah dari versi v0 atau v1. keduanya bisa saja dibangun secara bersamaan. Jadi, jika modul v2+ Anda memiliki beberapa paket, Anda harus mengubahnya menggunakan path impor yang baru "/v2". kalau tidak, modul v2+ Anda akan tetap bergantung pada modul v0 atau v1. Misalnya, untuk mengubah semua `example.com/hello` supaya mengacu ke `example/hello/v2`