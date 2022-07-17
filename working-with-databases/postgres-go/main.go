package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/santekno/postgres-go/services"
)

type app struct {
	AlbumService services.AlbumService
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := openDB(os.Getenv("POSTGRES_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)

	application := app{AlbumService: services.NewPostgresService(db)}

	albums, err := application.AlbumService.GetAllAlbum()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("all album : %v\n", albums)

	albumNo1, err := application.AlbumService.Get(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("album number 1 : %v\n", albumNo1)

	err = application.AlbumService.BatchCreate([]services.Album{
		{Title: "Hari Yang Cerah", Artist: "Peterpan", Price: 50000},
		{Title: "Sebuah Nama Sebuah Cerita", Artist: "Peterpan", Price: 50000},
		{Title: "Bintang Di surga", Artist: "Peterpan", Price: 60000},
	})
	if err != nil {
		log.Fatal(err)
	}

	albumNo1.Price = 70000
	err = application.AlbumService.Update(*albumNo1)
	if err != nil {
		log.Fatal(err)
	}

	err = application.AlbumService.Delete(albumNo1.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func openDB(dsn string, setLimits bool) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if setLimits {
		fmt.Println("setting limits")
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
