package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Mahasiswa struct {
	ID           int64
	Nama         string
	NIM          string
	JenisKelamin int
	TempatLahir  string
	TanggalLahir time.Time
	TahunMasuk   int
}

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
