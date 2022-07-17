package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/santekno/mssql-go/services"
)

type app struct {
	AlbumService services.AlbumService
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := openDB(os.Getenv("MSSQL_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)

	application := app{AlbumService: services.NewMSSQLService(db)}

	err = application.AlbumService.BatchCreate([]services.Album{
		{Title: "Hari Yang Cerah", Artist: "Peterpan", Price: 50000},
		{Title: "Sebuah Nama Sebuah Cerita", Artist: "Peterpan", Price: 50000},
		{Title: "Bintang Di surga", Artist: "Peterpan", Price: 60000},
	})
	if err != nil {
		log.Fatal(err)
	}

	albums, err := application.AlbumService.GetAllAlbum()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("all album : %v\n", albums)

	albumNo1, err := application.AlbumService.Get(albums[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("album number 1 : %v\n", albumNo1)

	err = application.AlbumService.Create(&services.Album{
		Title: "Mungkin Nanti", Artist: "Peterpan", Price: 70000,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success Create Album!")

	albumNo1.Price = 70000
	err = application.AlbumService.Update(*albumNo1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Success update album %v\n", albumNo1)

	err = application.AlbumService.Delete(albumNo1.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Success delete id: %d!\n", albumNo1.ID)

	for _, alb := range albums {
		err = application.AlbumService.Delete(alb.ID)
		log.Printf("error : %v", err)
	}
	log.Println("Success delete all data from table album")
}

func openDB(connString string, setLimits bool) (*sql.DB, error) {
	var err error
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	if setLimits {
		fmt.Println("setting limits")
		db.SetMaxIdleConns(5)
		db.SetMaxOpenConns(5)
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
	return db, nil
}
