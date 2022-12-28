package config

import (
	"database/sql"

	albumRepository "final-project-api/internal/repository/album"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	AlbumRepository albumRepository.AlbumRepository
}

// Function to initialize repository
func InitRepository(db *sql.DB, cache *redis.Client) Repository {
	return Repository{
		AlbumRepository: albumRepository.NewAlbumRepository(db, cache),
	}
}
