package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// albums, err := albumsByArtist("Peterpan")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Albums found: %v\n", albums)

	// for _, album := range albums {
	// 	log.Printf("Title  : %s\n", album.Title)
	// 	log.Printf("Artist : %s\n", album.Artist)
	// 	log.Printf("Price  : %f\n", album.Price)
	// }

	// Hard-code ID 2 here to test the query.
	// alb, err := albumByID(2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Album found: %v\n", alb)

	// albID, sumAffected, err := addAlbum(Album{
	// 	Title:  "Taman Langit",
	// 	Artist: "Peterpan",
	// 	Price:  50000.0,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ID of added album: %v\n", albID)
	// fmt.Printf("Sum affected %d\n", sumAffected)

	// // insert data bulk using transaction model
	// albumIDs, err := BulkInsertUsingTransaction()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ID of added album: %v\n", albumIDs)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func albumByID(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func addAlbum(alb Album) (int64, int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, 0, fmt.Errorf("addAlbum: %v", err)
	}

	sum, err := result.RowsAffected()
	if err != nil {
		return 0, 0, fmt.Errorf("error when getting rows affected")
	}

	return id, sum, nil
}

func BulkInsertUsingTransaction() ([]int64, error) {
	var insertID []int64

	tx, err := db.Begin()
	if err != nil {
		return insertID, fmt.Errorf("album transaction error")
	}

	defer tx.Rollback()

	albums := []Album{
		{
			Title:  "Hari Yang Cerah",
			Artist: "Peterpan",
			Price:  50000,
		},
		{
			Title:  "Sebuah Nama Sebuah Cerita",
			Artist: "Peterpan",
			Price:  50000,
		},
		{
			Title:  "Bintang Di surga",
			Artist: "Peterpan",
			Price:  60000,
		},
	}

	for _, album := range albums {
		result, err := tx.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
		if err != nil {
			log.Printf("error : %v", err)
			continue
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			log.Printf("error : %v", err)
		}

		insertID = append(insertID, lastInsertId)
		return insertID, errors.New("error")
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("error : %v", err)
		return insertID, err
	}

	return insertID, err
}
