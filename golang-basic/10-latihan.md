# Latihan

## Soal 1
Buat program `palindrome` menggunakan golang
```bash
$ go main.go 
$ madam
```
output
```bash
true
```

## Soal 2
Buat program untuk menentukan Huruf yang diinput Vokal atau Konsonan
```bash
$ go main .go
h
```
output
```bash
Konsonan
```

## Cara Mengumpulkan Latihan Dengan Pull Request
* clone repository ini dengan perintah seperti dibawah
  ```bash
  git clone git@github.com:santekno/tutorial-go.git
  ```
  Jika sudah selesai maka akan terlihat seperti ini
  ```bash
  $ git clone git@github.com:santekno/tutorial-go.git
  Cloning into 'tutorial-go'...
  remote: Enumerating objects: 133, done.
  remote: Counting objects: 100% (133/133), done.
  remote: Compressing objects: 100% (102/102), done.
  remote: Total 133 (delta 26), reused 130 (delta 23), pack-reused 0
  Receiving objects: 100% (133/133), 42.01 KiB | 177.00 KiB/s, done.
  Resolving deltas: 100% (26/26), done.
  ```
* Masuk ke folder sesi satu yaitu `golang-basic` dan lalu cari folder `latihan`.
  ```bash
  $ cd tutorial-go/golang-basic/latihan
  ```
* Kerjakan latihan ini sesuai dengan perintah diatas setelah selesai simpan file `go` tersebut di dalam folder `tutorial-golang/golang-basic/latihan` dengan aturan penamaan:
  ```bash
  <nomor-soal>-<nama-peserta>.go
  contoh:
  01-ihsan-arif.go
  02-ihsan-arif.go
  ```
* Jika latihan sudah dikerjakan dan disimpan pada folder `latihan` maka saatnya kita membuat `branch` di dalam repository ini.
  ```bash
  $ git checkout -b <nama-peserta>
  contoh: 
  $ git checkout -b ihsan-arif
  ```
* Lalu `commit` semua latihan ke github dengan perintah
  ```bash
  $ git push origin <nama-peserta>
  ```
* Jika sudah commit ke github maka saatnya kita buat PR (Pull Request). Buka repository lagi `github.com/santekno/tutorial-go` lalu pilih `Pull Requests`. Lalu `New Pull Request`.
* Setelah itu kita akan memilih branch mana yang akan kita `compare` dengan `main` branch. Kita pilih saja branch `<nama-peserta>`. `base:main` <- `compare:<nama-peserta>`.
* Isi Judul seperti ini `Latihan <nama-peserta> Golang Basic`
* Pilih tombol `Create Pull Requte`
* Selesai

## Menggunakan Forks
Jika teman-teman sudah terbiasa menggunakan `github` dengan mekanisme `forks`, boleh saja silahkan asalkan nanti tetap untuk membuat `Pull Requests` ke dalam main repository ya.

Terima kasih dan selamat mengerjakan.
