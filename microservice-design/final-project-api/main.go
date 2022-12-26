package main

import (
	"fmt"
	"log"
	"os"

	"final-project-api/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Initialize gin
	r := gin.Default()

	// Get the environment
	env := os.Getenv("ENV")
	if env != "production" {
		// Get the config from .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")

	// Load db config
	db, err := config.OpenDB(os.Getenv("POSTGRES_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer config.CloseDB(db)

	// Load redis
	cache := config.OpenCache(os.Getenv("REDIS"), "")

	// init grpc client
	grpcUrl := os.Getenv("GRPC_ARTIST_URL")
	artistGrpc := config.ServiceGrpcArtist(grpcUrl)

	// Init clean arch
	repository := config.InitRepository(db, cache, artistGrpc)
	usecase := config.InitUsecase(repository.AlbumRepository, repository.ArtistRepository)
	handler := config.InitHandler(usecase.AlbumUsecase, usecase.ArtistUsecase)

	// Create the API
	albumRoutes := r.Group("/api/v1/albums")
	{
		albumRoutes.GET("/", handler.AlbumHandler.GetAllAlbum)
		albumRoutes.POST("/", handler.AlbumHandler.Create)
		albumRoutes.POST("/batch", handler.AlbumHandler.BatchCreate)
		albumRoutes.GET("/:id", handler.AlbumHandler.Get)
		albumRoutes.PUT("/:id", handler.AlbumHandler.Update)
		albumRoutes.DELETE("/:id", handler.AlbumHandler.Delete)
	}

	// Create the API
	artistRoutes := r.Group("/api/v1/artists")
	{
		artistRoutes.GET("/", handler.ArtistHandler.GetAll)
		artistRoutes.GET("/:id", handler.ArtistHandler.Get)
	}

	// Run the gin gonic in port 5000
	runWithPort := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(runWithPort)
}
