package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/santekno/mysql-native/services"
)

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "students",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	service := services.InitServices(db)

	/*
	* BulkInsertMahasiswa
	*
	 */
	mahasiswaBulk := []services.Mahasiswa{
		{
			Nama:         "Ihsan",
			NIM:          "ABC001",
			JenisKelamin: 1,
			TempatLahir:  "Tasikmalaya",
			TanggalLahir: time.Date(1992, 10, 2, 0, 0, 0, 0, time.Local),
			TahunMasuk:   2020,
		},
		{
			Nama:         "Arif",
			NIM:          "ABC002",
			JenisKelamin: 1,
			TempatLahir:  "Tasikmalaya",
			TanggalLahir: time.Date(1992, 10, 3, 0, 0, 0, 0, time.Local),
			TahunMasuk:   2020,
		},
		{
			Nama:         "Rahman",
			NIM:          "ABC003",
			JenisKelamin: 1,
			TempatLahir:  "Tasikmalaya",
			TanggalLahir: time.Date(1992, 10, 4, 0, 0, 0, 0, time.Local),
			TahunMasuk:   2020,
		},
	}

	mahasiswaIDs, err := service.BulkInsertUsingTransaction(mahasiswaBulk)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of added mahasiswa: %v\n", mahasiswaIDs)

	/*
	* Get MahasiswaById
	* hardcoded id = 2
	 */
	mhs, err := service.GetMahasiswaById(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Mahasiswa found: %v\n", mhs)

	/*
	* Get All Mahasiswa
	*
	 */
	mahasiswas, err := service.GetAllMahasiswa()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mahasiswas found: %v\n", mahasiswas)

	for _, mhs := range mahasiswas {
		log.Printf("Nama  : %s\n", mhs.Nama)
		log.Printf("NIM : %s\n", mhs.NIM)
		log.Printf("Jenis Kelamin  : %d\n", mhs.JenisKelamin)
		log.Printf("Tempat Lahir  : %s\n", mhs.TempatLahir)
		log.Printf("Tanggal lahir : %v\n", mhs.TanggalLahir)
		log.Printf("Tahun Masuk : %d\n", mhs.TahunMasuk)
	}

	err = service.DeleteMahasiswa(mahasiswas[0].ID)
	if err != nil {
		log.Fatal(err)
	}
}
