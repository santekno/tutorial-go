# Cara Melakukan Instalasi Golang

Download terlebih dahulu golang library yang ada di website [Golang](https://golang.org/doc/install).

## MAC
* Pilih package dan download sesuai dengan MacOS support. Jika menggunakan Macbook M1 sebaiknya install yang versi ARM64.
* Setelah di download file berekstensi .dmg, lakukan install seperti biasanya melakukan instal aplikasi. Setelah selesai, biasanya path golang akan disimpan di folder `/usr/local/go`
* Jika dibutuhkan ketik perintah dibawah ini untuk melakukan set environment local
  ```bash
  export PATH=$PATH:/usr/local/go/bin
  ```

## Windows
* Pilih package dan download sesuai dengan versi Windows yang mendukung.
* Setelah download file berekstensi .exe, lakukan install seperti melakukan instalasi aplikasi lain. Biasanya file akan disimpan di folder `Program File` atau `Program File (x86)`
* Buka terminal atau `cmd`.

## Linux
* Pilih package dan download sesuai dengan versi Linux yang mendukung
* lakukan perintah dibawah ini
  ```bash
  rm -rf /usr/local/go && tar -C /usr/local -xzf go{version-go}.linux-amd64.tar.gz
  ```
* Pastikan path environment sudah di-set agar go library mendukung pada sistem operasi. Bisa dilakukan dengan perintah dibawah ini.
  ```bash
  export PATH=$PATH:/usr/local/go/bin
  ```

## Verifikasi
Setelah melakukan proses instalasi tersebut diatas kita bisa mengecek apakah Golang ini sudah terinstal di komputer kita atau belum dengan cara 
* Buka `terminal`
* lalu ketik perintah `go version`, maka jika sudah terinstall akan keluar informasi seperti ini
  ```bash
  go version go1.17.5 darwin/arm64
  ```
* atau jika kamu tahu lebih banyak terkait syntax yang lebih detail bisa pakai `go help`.
* 