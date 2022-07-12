package main

import (
	"log"

	"github.com/santekno/using-gorm/internal/album"
	"github.com/santekno/using-gorm/internal/album/repository"
	"github.com/santekno/using-gorm/internal/database"
	"github.com/santekno/using-gorm/internal/domain"
)

type App struct {
	AlbumRepo album.AlbumRepo
}

func main() {
	db := database.NewDatabaseConn()

	app := App{AlbumRepo: repository.NewAlbumRepo(db)}

	albums, err := app.AlbumRepo.GetAllAlbum()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("success get all data album : %+v", albums)

	album, err := app.AlbumRepo.GetAlbumById(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("success get by id 1 : %+v", album)

	err = app.AlbumRepo.CreateAlbum(domain.Album{ID: 5, Title: "Taman langit", Artist: "Peterpan", Price: 35000})
	if err != nil {
		log.Fatal(err)
	}

	albumUpdate := []domain.Album{
		{
			ID:     1,
			Title:  "Hari Yang Cerah",
			Artist: "Peterpan",
			Price:  50000,
		},
		{
			ID:     2,
			Title:  "Sebuah Nama Sebuah Cerita",
			Artist: "Peterpan",
			Price:  50000,
		},
		{
			ID:     3,
			Title:  "Bintang Di surga",
			Artist: "Peterpan",
			Price:  60000,
		},
	}
	// bulk insert
	err = app.AlbumRepo.BulkInsertAlbum(albumUpdate)
	if err != nil {
		log.Fatal(err)
	}

	// update
	err = app.AlbumRepo.UpdateAlbum(domain.Album{ID: 3, Title: "Taman langit", Artist: "Peterpan", Price: 35000})
	if err != nil {
		log.Fatal(err)
	}

	err = app.AlbumRepo.DeleteAlbum(3)
	if err != nil {
		log.Fatal(err)
	}
}
