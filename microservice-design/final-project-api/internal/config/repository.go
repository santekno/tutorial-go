package config

import (
	"database/sql"

	albumRepository "final-project-api/internal/repository/album"
	artistRepository "final-project-api/internal/repository/artist"
	"final-project-api/internal/repository/artist/grpc"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	AlbumRepository  albumRepository.AlbumRepository
	ArtistRepository artistRepository.ArtistRepository
}

// Function to initialize repository
func InitRepository(db *sql.DB, cache *redis.Client, client grpc.ArtistInterface) Repository {
	return Repository{
		AlbumRepository:  albumRepository.NewAlbumRepository(db, cache),
		ArtistRepository: artistRepository.NewArtistRepository(client),
	}
}
